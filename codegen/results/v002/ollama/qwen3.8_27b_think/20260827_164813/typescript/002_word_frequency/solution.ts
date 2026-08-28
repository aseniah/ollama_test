import * as fs from 'fs';

const text = fs.readFileSync('input/text.txt', 'utf-8');

// Split into words, lowercase, strip non-letter characters
const words = text
  .toLowerCase()
  .split(/\s+/)
  .map(w => w.replace(/[^a-z]/g, ''))
  .filter(w => w.length > 0);

// Count frequencies
const freq: Map<string, number> = new Map();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array and sort: count descending, then alphabetically ascending
const entries = Array.from(freq.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
process.stdout.write(output + '\n');