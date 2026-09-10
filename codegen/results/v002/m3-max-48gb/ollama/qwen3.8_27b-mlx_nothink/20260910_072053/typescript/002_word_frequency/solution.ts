import { readFileSync } from 'fs';

// Read the file
const text = readFileSync('input/text.txt', 'utf-8');

// Extract words: keep only letters, convert to lowercase
// Split on non-letter characters to get words
const words = text
  .toLowerCase()
  .split(/[^a-z]+/)  // split on non-lowercase-letter sequences
  .filter(word => word.length > 0);  // filter out empty strings

// Count frequency of each word
const wordCounts = new Map<string, number>();
for (const word of words) {
  wordCounts.set(word, (wordCounts.get(word) || 0) + 1);
}

// Convert to array of [word, count] pairs
const entries = Array.from(wordCounts.entries());

// Sort by count descending, then alphabetically ascending for ties
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output results
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}