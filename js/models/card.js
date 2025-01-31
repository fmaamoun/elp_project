export default class Card {
    constructor(id, words) {
      this.id = id;
      this.words = words; // Array of 5 words
    }
  
    // Selects a random word from the card
    getRandomWord() {
      const index = Math.floor(Math.random() * this.words.length);
      return this.words[index];
    }
  }
  