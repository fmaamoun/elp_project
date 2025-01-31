import readline from 'readline-sync';
import fs from 'fs';
import path from 'path';
import Player from './Player.js';
import { shuffleArray } from '../utils/helpers.js';
import Card from './Card.js';

export default class Game {
  constructor(cards) {
    // Convert card data objects into Card instances
    this.allCards = cards.map(cardData => new Card(cardData.id, [
      cardData.word1,
      cardData.word2,
      cardData.word3,
      cardData.word4,
      cardData.word5
    ]));
    this.usedCards = [];
    this.players = [];
    this.propositionsLogPath = path.join('logs', 'propositions.log');
    // Initialize or clear the log file
    fs.writeFileSync(this.propositionsLogPath, '', 'utf8');
  }

  // Initializes player names
  initializePlayers() {
    const numberOfPlayers = 5;
    console.log(`\nPlease enter the names of the ${numberOfPlayers} players:`);
    for (let i = 1; i <= numberOfPlayers; i++) {
      let name;
      do {
        name = readline.question(`Player ${i} name: `).trim();
        if (name === '') {
          console.log("Name cannot be empty. Please try again.");
        }
      } while (name === '');
      this.players.push(new Player(name));
    }
  }

  // Starts the game and manages the game flow
  start() {
    // Shuffle the cards
    this.cards = shuffleArray([...this.allCards]);
    const numberOfRounds = Math.min(this.cards.length, 10); // Example: 10 rounds

    for (let round = 1; round <= numberOfRounds; round++) {
      console.log(`\n--- Round ${round} ---`);

      // Select the guesser
      const guesserIndex = (round - 1) % this.players.length;
      const guesser = this.players[guesserIndex];
      console.log(`The guesser is: ${guesser.name}`);

      // Draw a card
      if (this.cards.length === 0) {
        console.log("No more cards available.");
        break;
      }
      const card = this.cards.shift();
      this.usedCards.push(card);

      // Select a random secret word from the card
      const secretWord = card.getRandomWord();

      // Instructions to players
      console.log("\nClue-givers, please look at the secret word.");
      console.log("Guesser, please look away.\n");
      
      // Display the secret word to clue-givers
      console.log(`Secret Word (visible to clue-givers): ${secretWord}\n`);

      // Pause to allow clue-givers to read the secret word
      readline.question("Press Enter when ready to enter proposals...");

      // Clear the console to hide the secret word from the guesser
      console.clear();

      // Collect proposals from other players
      const proposals = [];
      for (let i = 0; i < this.players.length; i++) {
        if (i === guesserIndex) continue; // Skip the guesser
        const player = this.players[i];
        let proposal;
        do {
          proposal = readline.question(`${player.name}, enter your one-word proposal: `).trim();
          if (proposal === '') {
            console.log("Proposal cannot be empty. Please try again.");
          }
        } while (proposal === '');
        proposals.push(proposal);

        // Log the proposal
        const logEntry = `Round ${round} - ${player.name}: ${proposal} (Card ID: ${card.id})\n`;
        fs.appendFileSync(this.propositionsLogPath, logEntry, 'utf8');
      }

      // Remove duplicate proposals (case-insensitive)
      const proposalsLower = proposals.map(p => p.toLowerCase());
      const duplicates = proposalsLower.filter((word, index, self) => self.indexOf(word) !== index);
      const uniqueProposals = proposals.filter((p, index) => !duplicates.includes(p.toLowerCase()));

      if (uniqueProposals.length === 0) {
        console.log("\nAll proposals were duplicates and have been removed.");
      } else {
        console.log("\nUnique proposals after removing duplicates:");
        uniqueProposals.forEach((word, index) => {
          console.log(`${index + 1}. ${word}`);
        });
      }

      // Guesser attempts to guess the secret word
      const attempt = readline.question(`\n${guesser.name}, enter your guess: `).trim();

      if (attempt.toLowerCase() === secretWord.toLowerCase()) {
        console.log("Congratulations! You guessed the word correctly.");
        guesser.score += 1;
      } else {
        console.log(`Sorry! The correct word was: ${secretWord}`);
      }
    }

    // Display final scores
    console.log("\n--- Final Scores ---");
    this.players.forEach(player => {
      console.log(`${player.name}: ${player.score} point(s)`);
    });
  }
}
