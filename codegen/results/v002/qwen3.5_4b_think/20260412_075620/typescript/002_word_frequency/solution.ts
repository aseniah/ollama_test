import { readFileSync, writeFileSync } from 'fs';

// Read the input file
const text = readFileSync('input/text.txt', 'utf8');

// Process text: convert to lowercase and extract words
const words = text
  .toLowerCase()
  .replace(/[^a-z0-9]/g, '') // Remove all characters that are not letters
  .split(/\s+/); // Split by whitespace

// Count frequency of each word
const frequency: Record<string, number> = {};
for (const word of words) {
  if (word) {
    frequency[word] = (frequency[word] || 0) + 1;
  }
}

// Sort by count descending, then alphabetically ascending
const sortedWords = Object.entries(frequency)
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Sort by count descending
    }
    return a[0].localeCompare(b[0]); // Sort alphabetically ascending for ties
  });

// Output results
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}