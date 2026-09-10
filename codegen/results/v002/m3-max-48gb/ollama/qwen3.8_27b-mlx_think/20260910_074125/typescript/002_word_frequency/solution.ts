import * as fs from 'fs';

const text = fs.readFileSync('input/text.txt', 'utf-8');

// Convert to lowercase and extract only words (sequences of letters)
const words = text.toLowerCase().match(/[a-z]+/g) || [];

// Count frequency
const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Sort by count descending, then alphabetically ascending for ties
const sorted = [...freq.entries()].sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
const lines = sorted.map(([word, count]) => `${word}: ${count}`);
process.stdout.write(lines.join('\n') + '\n');