import fs from "fs"
import csv from "csv-parser"

export function loadCards(filePath) {
  return new Promise((resolve, reject) => {
    const cards = []
    fs.createReadStream(filePath)
      .pipe(csv())
      .on("data", (row) => {
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
        resolve(cards)
      })
      .on("error", (error) => {
        reject(error)
      })
  })
}
