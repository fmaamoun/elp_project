import fs from "fs"
import csv from "csv-parser"

/**
 * Reads a CSV file located at filePath and returns a Promise that resolves with an array of card objects.
 * Each card object will include an `id` and 5 words (`word1` through `word5`).
 *
 */
export function loadCards(filePath) {
  return new Promise((resolve, reject) => {
    const cards = []

    // Create a readable stream from the CSV file, pipe it through csv-parser,
    // and build an array of card objects.
    fs.createReadStream(filePath)
      .pipe(csv())
      .on("data", (row) => {
        // Each row corresponds to a line in the CSV file
        cards.push({
          id: parseInt(row.id, 10),
          word1: row.word1.trim(),
          word2: row.word2.trim(),
          word3: row.word3.trim(),
          word4: row.word4.trim(),
          word5: row.word5.trim(),
        })
      })
      .on("end", () => {
        // Once the stream is done, resolve the Promise with the cards array
        resolve(cards)
      })
      .on("error", (error) => {
        // If there's any error during reading/parsing, reject the Promise
        reject(error)
      })
  })
}
