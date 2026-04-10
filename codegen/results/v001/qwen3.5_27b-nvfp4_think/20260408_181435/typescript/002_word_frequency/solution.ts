import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Split into words, keep only letters, convert to lowercase
const words = content
  .split(/\s+/) // split on whitespace
  .map(word => word.toLowerCase())
  .filter(word => word.length > 0);

const wordCounts: Record<string, number> = {};

for (const word of words) {
  const cleanWord = word.replace(/[^a-z]/g, ''); // keep only letters
  if (cleanWord.length > 0) {
    wordCounts[cleanWord] = (wordCounts[cleanWord] || 0) + 1;
  }
}

// Convert to array and sort
const sortedWords = Object.entries(wordCounts).sort(
  (a, b) => b[1] - a[1] || a[0].localeCompare(b[0])
);

// Output
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}