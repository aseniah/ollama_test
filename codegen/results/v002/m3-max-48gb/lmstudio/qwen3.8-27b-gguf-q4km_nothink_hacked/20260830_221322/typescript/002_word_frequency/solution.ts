import { readFileSync } from 'fs';

const content = readFileSync('input/text.txt', 'utf-8');

// Convert to lowercase
const lower = content.toLowerCase();

// Keep only letters and spaces (strip all punctuation)
const cleaned = lower.replace(/[^a-z\s]/g, '');

// Split into words by whitespace
const words = cleaned.split(/\s+/).filter(w => w.length > 0);

// Count frequency
const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array of [word, count]
const entries = Array.from(freq.entries());

// Sort by count descending, then alphabetically ascending
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}