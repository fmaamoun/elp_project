import path from 'path';
import { fileURLToPath } from 'url';
import { loadCards } from './utils/fileHandler.js';
import Game from './models/Game.js';

// Helper to get the directory name in ES modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

async function main() {
  console.log("Welcome to 'Just One'!");

  // Load cards from the CSV file
  const cardsPath = path.join(__dirname, '..', 'js/data', 'cards.csv');
  let cards;
  try {
    cards = await loadCards(cardsPath);
    if (cards.length === 0) {
      throw new Error("No cards loaded. Please check the CSV file.");
    }
  } catch (error) {
    console.error("Error loading cards:", error.message);
    process.exit(1);
  }

  // Initialize the game
  const game = new Game(cards);
  game.initializePlayers();

  // Start the game
  game.start();

  console.log("\nThank you for playing 'Just One'!");
  console.log(`All proposals have been saved to 'logs/propositions.log'.`);
}

main();
