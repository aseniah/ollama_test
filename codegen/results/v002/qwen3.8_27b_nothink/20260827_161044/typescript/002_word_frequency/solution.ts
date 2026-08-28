import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'text.txt');
const text = fs.readFileSync(inputPath, 'utf-8');

// Convert to lowercase, keep only letters (strip punctuation)
const words = text
  .toLowerCase()
  .replace(/[^a-z\s]/g, ' ')  // replace non-letter, non-whitespace with space
  .split(/\s+/)
  .filter(w => w.length > 0);

const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array and sort
const entries = Array.from(freq.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  // Alphabetical ascending
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}