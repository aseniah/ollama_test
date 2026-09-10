import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.resolve('input/text.txt');
const content = fs.readFileSync(inputPath, 'utf-8');

// Split into tokens by whitespace, strip non-letters, lowercase
const words = content
  .split(/\s+/)
  .map((word) => word.replace(/[^a-z]/gi, '').toLowerCase())
  .filter((word) => word.length > 0);

// Count frequency
const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Sort by count descending, then alphabetically ascending
const sorted = Array.from(freq.entries()).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}