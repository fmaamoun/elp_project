/**
 * Represents a single card in the game, containing an ID and an array of 5 words.
 */
export default class Card {
  constructor(id, words) {
    this.id = id
    this.words = words // Array of 5 words
  }
}
