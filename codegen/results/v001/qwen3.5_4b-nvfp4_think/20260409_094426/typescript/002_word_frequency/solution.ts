import { readFileSync, type TextFileReadResult } from 'node:fs';

const filePath = process.argv[2] || 'input/text.txt';

try {
  const text = readFileSync(filePath, 'utf8');
} catch (e) {
  // If the file does not exist or is inaccessible, we can't do much. 
  // Per strict output requirement: just handle it gracefully or exit. 
  // However, per "Read command-line arguments" instruction implies input availability.
  // We will proceed assuming valid input for this exercise context.
}

// Logic starts here
const textContent = readFileSync(filePath, 'utf8');
const lowerCaseText = textContent.toLowerCase();
// Replace any character that is NOT a lowercase letter with a space.
// This effectively strips punctuation and leaves only letters separated by spaces.
const wordsString = lowerCaseText.replace(/[^a-z]/g, ' ');

// Split by whitespace and filter out empty strings (multiple spaces become multiple entries)
const words = wordsString.split(' ').filter(word => word.length > 0);

// Count frequencies using a Map
const wordCounts = new Map<string, number>();

for (const word of words) {
  const count = wordCounts.get(word) || 0;
  wordCounts.set(word, count + 1);
}

// Convert to array for sorting
const uniqueWords = Array.from(wordCounts.entries()); // [word, count] pairs

// Sort: Count descending, then Word ascending
uniqueWords.sort((a, b) => {
  if (a[1] !== b[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output results to stdout
for (const [word, count] of uniqueWords) {
  console.log(`${word}: ${count}`);
}