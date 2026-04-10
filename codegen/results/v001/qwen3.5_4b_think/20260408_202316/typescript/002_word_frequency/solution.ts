import * as fs from 'fs';
import * as path from 'path';

// Default file path or use command-line argument
const filePath = process.argv[2] || 'input/text.txt';

// Read file content
const content = fs.readFileSync(filePath, 'utf-8');

// Extract words: convert to lowercase, keep only letters
const words = content
  .toLowerCase()
  .match(/[a-z]+/g)
  || [];

// Count word frequencies
const frequencyMap: Map<string, number> = new Map();
for (const word of words) {
  if (word.length > 0) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }
}

// Sort by count descending, then alphabetically ascending
const sorted = Array.from(frequencyMap.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
  });

// Output results
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}