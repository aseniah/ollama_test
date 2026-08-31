import { readFileSync } from 'fs';

// Read the file
const text = readFileSync('input/text.txt', 'utf-8');

// Split into words, convert to lowercase, strip non-letter characters
const words = text
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequencies
const freq: Map<string, number> = new Map();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array and sort
const entries = Array.from(freq.entries());
entries.sort((a, b) => {
  // Sort by count descending
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  // For ties, sort alphabetically ascending
  return a[0].localeCompare(b[0]);
});

// Output
const lines = entries.map(([word, count]) => `${word}: ${count}`);
process.stdout.write(lines.join('\n') + '\n');