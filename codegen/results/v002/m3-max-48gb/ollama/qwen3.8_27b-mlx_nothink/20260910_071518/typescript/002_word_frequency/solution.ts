import * as fs from 'fs';

// Read the file
const content = fs.readFileSync('input/text.txt', 'utf-8');

// Convert to lowercase and keep only letters (a-z) and spaces
// Replace all non-letter, non-space characters with space, then collapse spaces
const normalized = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, ' ')
  .replace(/\s+/g, ' ')
  .trim();

// Split into words
const words = normalized.split(' ');

// Count frequency
const freq: Record<string, number> = {};
for (const word of words) {
  if (word.length > 0) {
    freq[word] = (freq[word] || 0) + 1;
  }
}

// Sort by count descending, then alphabetically ascending
const entries = Object.entries(freq);
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}