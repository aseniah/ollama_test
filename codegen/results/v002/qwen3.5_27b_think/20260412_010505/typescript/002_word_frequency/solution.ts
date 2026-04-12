import * as fs from 'fs';

const filePath = process.argv[2] || 'input/text.txt';
const text = fs.readFileSync(filePath, 'utf-8');

// Convert to lowercase
const lowerText = text.toLowerCase();

// Extract words (only letters)
const words = lowerText.match(/[a-z]+/g) || [];

// Count frequencies
const frequency = new Map<string, number>();
for (const word of words) {
  frequency.set(word, (frequency.get(word) || 0) + 1);
}

// Convert to array and sort
const entries = Array.from(frequency.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1]; // Descending by count
  }
  return a[0].localeCompare(b[0]); // Ascending alphabetically
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}