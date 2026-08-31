import { readFileSync } from 'fs';

const text = readFileSync('input/text.txt', 'utf-8');
const lower = text.toLowerCase();

// Extract all words (sequences of letters)
const words = lower.match(/[a-z]+/g) || [];

const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

const entries = Array.from(freq.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}