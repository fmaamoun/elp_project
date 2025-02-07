import readline from "readline-sync"
import fs from "fs"
import Player from "./player.js"
import { shuffleArray } from "../utils/helpers.js"
import Card from "./Card.js"

export default class Game {
  constructor(cards) {
    // Convert card data objects into Card instances
    this.allCards = cards.map(
      (cardData) =>
        new Card(cardData.id, [
          cardData.word1,
          cardData.word2,
          cardData.word3,
          cardData.word4,
          cardData.word5,
        ]),
    )

    this.usedCards = []
    this.players = []

    // Global team score
    this.teamScore = 0

    // Name of the file that will store proposals
    this.proposalsFile = "proposals.log"
  }

  // Prompts the user to input player names
  initializePlayers() {
    const numberOfPlayers = 5
    console.log(`\nPlease enter the names of the ${numberOfPlayers} players:`)

    for (let i = 1; i <= numberOfPlayers; i++) {
      let name
      do {
        name = readline.question(`Player ${i} name: `).trim()
        if (name === "") {
          console.log("Name cannot be empty. Please try again.")
        }
      } while (name === "")
      this.players.push(new Player(name))
    }
  }

  /**
   * MAIN entry point for the game.
   * 1) Removes existing proposal file (if any).
   * 2) Shuffles cards and picks 13 of them.
   * 3) Loops through rounds while there are cards left.
   * 4) Displays final score and message at the end.
   */
  start() {
    this.removeExistingProposalsFile()
    this.setupCards()

    let round = 1

    // Loop as long as there are cards left
    while (this.cards.length > 0) {
      this.playRound(round)
      round++
    }

    // When no more cards, show final score
    this.displayFinalScore()
  }

  /**
   * Removes the proposals file (if it exists) to start fresh.
   */
  removeExistingProposalsFile() {
    if (fs.existsSync(this.proposalsFile)) {
      fs.unlinkSync(this.proposalsFile)
    }
  }

  /**
   * Shuffles all cards and slices the deck to 13 cards.
   */
  setupCards() {
    this.cards = shuffleArray([...this.allCards]).slice(0, 13)
  }

  /**
   * Handles a single round of the game:
   */
  playRound(round) {
    console.log(`\n--- Round ${round} ---`)

    // Identify the guesser (based on round number)
    const guesserIndex = (round - 1) % this.players.length
    const guesser = this.players[guesserIndex]
    console.log(`The guesser is: ${guesser.name}`)

    // Safety check in case the deck is empty
    if (this.cards.length === 0) {
      console.log("No more cards available.")
      return
    }

    // Look at the top card (but don't remove it yet)
    const card = this.cards[0]

    // Ask guesser which word index (1-5)
    const selectedWordIndex = this.getSelectedWordIndex(guesser)
    const secretWord = card.words[selectedWordIndex - 1]

    // Show secret word to clue-givers, let them read it
    this.showSecretWordToClueGivers(secretWord)

    // Collect proposals from all clue-givers (except guesser)
    const proposals = this.collectProposals(
      guesserIndex,
      selectedWordIndex,
      round,
    )

    // Remove duplicates among proposals
    const uniqueProposals = this.removeDuplicateProposals(proposals)

    // Display unique proposals (or note if all were duplicates)
    this.displayProposals(uniqueProposals)

    // Handle the guess (or passing)
    this.handleGuess(guesser, secretWord)
  }

  /**
   * Prompts the guesser for a valid index (1-5).
   */
  getSelectedWordIndex(guesser) {
    let selectedWordIndex
    do {
      selectedWordIndex = parseInt(
        readline.question(
          `\n${guesser.name}, enter the number (1-5) of the word you want to guess: `,
        ),
      )
    } while (
      isNaN(selectedWordIndex) ||
      selectedWordIndex < 1 ||
      selectedWordIndex > 5
    )
    return selectedWordIndex
  }

  /**
   * Displays the secret word to clue-givers, prompts them to proceed,
   * and then clears the console so the guesser cannot see it.
   */
  showSecretWordToClueGivers(secretWord) {
    console.log("\nClue-givers, please look at the secret word.")
    readline.question("Guesser, please look away.\n")

    console.log(`Secret Word (visible to clue-givers): ${secretWord}\n`)
    readline.question("Press Enter when ready to enter proposals...")

    console.clear() // Hide the secret word from the guesser
  }

  /**
   * Asks each player (except guesser) for a one-word proposal,
   * and appends each proposal to the proposalsFile.
   */
  collectProposals(guesserIndex, selectedWordIndex, round) {
    const proposals = []

    for (let i = 0; i < this.players.length; i++) {
      if (i === guesserIndex) continue
      const player = this.players[i]

      let proposal
      do {
        proposal = readline
          .question(
            `${player.name}, enter your one-word proposal for word ${selectedWordIndex}: `,
          )
          .trim()
        if (proposal === "") {
          console.log("Proposal cannot be empty. Please try again.")
        }
      } while (proposal === "")

      proposals.push(proposal)

      // Immediately save each proposal to the file
      fs.appendFileSync(
        this.proposalsFile,
        `Round=${round}, ClueGiver=${player.name}, WordIndex=${selectedWordIndex}, Proposal=${proposal}\n`,
      )
    }

    return proposals
  }

  /**
   * Removes duplicate proposals based on case-insensitive comparison.
   */
  removeDuplicateProposals(proposals) {
    const proposalsLower = proposals.map((p) => p.toLowerCase())
    const duplicates = proposalsLower.filter(
      (word, index, self) => self.indexOf(word) !== index,
    )
    return proposals.filter((p) => !duplicates.includes(p.toLowerCase()))
  }

  /**
   * Displays unique proposals, or notes if there are none left.
   */
  displayProposals(uniqueProposals) {
    if (uniqueProposals.length === 0) {
      console.log("\nAll proposals were duplicates and have been removed.")
    } else {
      console.log("\nUnique proposals after removing duplicates:")
      uniqueProposals.forEach((word, index) => {
        console.log(`${index + 1}. ${word}`)
      })
    }
  }

  /**
   * Prompts the guesser to either pass or guess.
   */
  handleGuess(guesser, secretWord) {
    const attempt = readline
      .question(`\n${guesser.name}, enter your guess or type 'pass' to skip: `)
      .trim()

    // If guesser passes, remove 1 card from the top
    if (attempt.toLowerCase() === "pass") {
      console.log("You chose to pass. No points awarded.")
      this.cards.shift() // -1 card
      console.log(`Remaining cards in the deck: ${this.cards.length}`)
      return
    }

    // Otherwise, compare attempt to secretWord
    if (attempt.toLowerCase() === secretWord.toLowerCase()) {
      console.log("Congratulations! You guessed the word correctly.")
      this.teamScore += 1
      this.cards.shift() // -1 card
      console.log(`Remaining cards in the deck: ${this.cards.length}`)
    } else {
      console.log(`Sorry! The correct word was: ${secretWord}`)
      this.cards.shift() // remove current card
      if (this.cards.length > 0) {
        this.cards.shift() // remove extra card
      }
      console.log(`Remaining cards in the deck: ${this.cards.length}`)
    }
  }

  /**
   * Called once the deck is exhausted or no more rounds.
   * Displays the team's final score and a message based on performance.
   */
  displayFinalScore() {
    console.log("\n--- Final Score ---")
    console.log(`Team Score: ${this.teamScore}`)

    switch (true) {
      case this.teamScore === 13:
        console.log("Score parfait ! Y arriverez-vous encore ?")
        break
      case this.teamScore === 12:
        console.log("Incroyable! Vos amis doivent être impressionnés !")
        break
      case this.teamScore === 11:
        console.log("Génial! C'est un score qui se fête!")
        break
      case this.teamScore >= 9 && this.teamScore <= 10:
        console.log("Waouh, pas mal du tout !")
        break
      case this.teamScore >= 7 && this.teamScore <= 8:
        console.log("Vous êtes dans la moyenne. Arriverez-vous à faire mieux ?")
        break
      case this.teamScore >= 4 && this.teamScore <= 6:
        console.log("C'est un bon début. Réessayez !")
        break
      case this.teamScore >= 0 && this.teamScore <= 3:
        console.log("Essayez encore.")
        break
      default:
        console.log("Quelque chose d'étrange s'est produit avec le score...")
    }
  }
}
