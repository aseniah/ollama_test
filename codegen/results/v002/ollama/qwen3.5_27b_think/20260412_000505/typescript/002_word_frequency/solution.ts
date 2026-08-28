import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');

// Convert to lowercase, split by whitespace, keep only letters
const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequency
const frequency: Record<string, number> = {};
for (const word of words) {
  frequency[word] = (frequency[word] || 0) + 1;
}

// Convert to array and sort by count descending, then alphabetically ascending
const entries = Object.entries(frequency);
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output each word with its count
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}