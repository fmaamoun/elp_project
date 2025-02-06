
# JustOne

**JustOne** is a cooperative party game where players try to guess words with the help of one-word clues given by their teammates. This console-based version allows you to play the game right in your terminal.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Formatting with Prettier](#formatting-with-prettier)
## Features
- Prompts players to enter their names.
- Shuffles a set of cards (CSV-based) and deals 13 of them.
- Asks one player (the guesser) to choose which of the five words on a card to guess.
- The other players each propose a one-word clue (duplicates are removed).
- The guesser either guesses or passes:
  - **Pass**: Only one card is discarded.
  - **Correct Guess**: One card is discarded, and the team score increases by 1.
  - **Incorrect Guess**: Two cards are discarded (the current card plus one additional).
- Final team score is displayed with a fun message based on performance.
- All proposals from clue-givers are recorded in a text file (`proposals.txt`).

## Prerequisites
- **Node.js** (version 14 or higher recommended)
- **npm** (or **yarn**, if you prefer)

## Installation
1. **Clone** this repository to your local machine:
   ```bash
   git clone https://github.com/fmaamoun/elp_project/tree/main/js
   ```

2. **Install** dependencies:
   ```bash
   npm install
   ```
   or
   ```bash
   yarn
   ```

## Usage
1. Make sure your CSV file of cards (e.g., `cards.csv`) is placed in the `src/data` folder.
2. Run the game:
   ```bash
   npm start
   ```
   or
   ```bash
   node src/index.js
   ```
3. Follow the prompts in your console to play the game!

## File Structure
A simplified overview of the project layout:
```
.
├── src
│   ├── data
│   │   └── cards.csv
│   ├── models
│   │   ├── card.js
│   │   ├── game.js
│   │   └── player.js
│   ├── utils
│   │   ├── file-handler.js
│   │   └── helpers.js
│   └── index.js
├── package.json
├── package-lock.json (or yarn.lock)
└── README.md
```

- **`src/index.js`**: Entry point of the application.  
- **`src/models/`**: Contains class definitions for `Game`, `Card`, and `Player`.  
- **`src/utils/`**: Helpers like reading CSV files (`file-handler.js`) and shuffling arrays (`helpers.js`).  
- **`src/data/cards.csv`**: Example CSV data file containing the words used in the game.

## Formatting with Prettier
To format all files in this project with **Prettier**, run:
```bash
npx prettier . --write
```
This command will automatically fix code style issues in supported file types.
