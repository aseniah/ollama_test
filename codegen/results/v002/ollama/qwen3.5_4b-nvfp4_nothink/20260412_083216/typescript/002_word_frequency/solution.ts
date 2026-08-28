import { readFileSync } from "fs";

const textFile = process.argv[2] || "input/text.txt";
const content = readFileSync(textFile, "utf-8");

// Split into words and filter/normalize them
const words = content
  .toLowerCase() // Convert to lowercase
  .split(/[^a-z]+/) // Split by any non-letter character (spaces, punctuation, etc.)
  .map((w) => w.trim()) // Trim any whitespace from the split
  .filter((w) => w.length > 0) // Remove empty strings

// Object to store word frequencies
const counts = new Map<string, number>();

for (const word of words) {
  const currentCount = counts.get(word) || 0;
  counts.set(word, currentCount + 1);
}

// Sort the unique words by count (descending), then alphabetically (ascending)
const sortedWords = Array.from(counts.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Descending order by count
    } else {
      return a[0].localeCompare(b[0]); // Ascending order for ties
    }
  });

// Output the results in the format: word: count
sortedWords.forEach(([word, count]) => {
  console.log(`${word}: ${count}`);
});