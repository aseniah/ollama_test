import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Normalize to lowercase
const lowerContent = content.toLowerCase();

// Extract words: keep only sequences of letters (a-z)
const words = lowerContent.match(/[a-z]+/g) || [];

const wordCounts = new Map<string, number>();

for (const word of words) {
  const current = wordCounts.get(word) || 0;
  wordCounts.set(word, current + 1);
}

// Convert to array of [word, count]
const entries = Array.from(wordCounts.entries());

// Sort: count descending, then word ascending
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
console.log(output);