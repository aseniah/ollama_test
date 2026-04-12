import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const filePath = process.argv[2] || 'input/text.txt';
const text = fs.readFileSync(path.resolve(filePath), 'utf8');

// Split into words, convert to lowercase, strip non-letter characters
const words = text
  .toLowerCase()
  .split(/\s+/)
  .filter(word => word.length > 0)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequency of each word
const frequency: Record<string, number> = {};

for (const word of words) {
  frequency[word] = (frequency[word] || 0) + 1;
}

// Sort by count descending, then alphabetically ascending for ties
const sorted = Object.entries(frequency).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1]; // count descending
  }
  return a[0].localeCompare(b[0]); // alphabetically ascending for ties
});

// Output results
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}