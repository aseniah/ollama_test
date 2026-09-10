import * as fs from 'fs';

const text = fs.readFileSync('input/text.txt', 'utf-8');

// Split into words by whitespace, lowercase, strip non-letter chars
const words = text
  .toLowerCase()
  .split(/\s+/)
  .map(w => w.replace(/[^a-z]/g, ''))
  .filter(w => w.length > 0);

// Count frequencies
const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) ?? 0) + 1);
}

// Sort by count descending, then alphabetically ascending
const sorted = Array.from(freq.entries()).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
const output = sorted.map(([word, count]) => `${word}: ${count}`).join('\n');
process.stdout.write(output + '\n');