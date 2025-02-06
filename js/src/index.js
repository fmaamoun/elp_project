import path from "path"
import { fileURLToPath } from "url"
import { loadCards } from "./utils/fileHandler.js"
import Game from "./models/game.js"

/**
 * Helper to get the directory name in ES modules (similar to __dirname in CommonJS).
 */
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

/**
 * Main entry point of the application. Loads the cards, initializes the game, then starts it.
 */
async function main() {
  console.log("Welcome to 'Just One'!")

  // Load the CSV file containing the card data
  const cardsPath = path.join(__dirname, "..", "src/data", "cards.csv")
  let cards
  try {
    // Attempt to load cards from the CSV file
    cards = await loadCards(cardsPath)

    // If no cards were found, throw an error
    if (cards.length === 0) {
      throw new Error("No cards loaded. Please check the CSV file.")
    }
  } catch (error) {
    // If something goes wrong (file not found, parse error, etc.), log it and exit
    console.error("Error loading cards:", error.message)
    process.exit(1)
  }

  // Create a new Game instance using the loaded cards
  const game = new Game(cards)

  // Prompt the user to enter the players' names
  game.initializePlayers()

  // Start the game (rounds, guesses, etc.)
  game.start()

  // Display a closing message
  console.log("\nThank you for playing 'Just One'!")
}

// Invoke the main function
main()
