import { readFileSync } from 'fs';

const content = readFileSync('input/text.txt', 'utf-8');

// Extract words: keep only letters, lowercase
const words = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, '')
  .split(/\s+/)
  .filter(w => w.length > 0);

// Count frequency
const freq: Map<string, number> = new Map();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Sort by count descending, then alphabetically ascending
const sorted = Array.from(freq.entries()).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
const output = sorted.map(([word, count]) => `${word}: ${count}`).join('\n');
process.stdout.write(output + '\n');