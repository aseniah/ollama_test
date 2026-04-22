import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Convert to lowercase
const lowerContent = content.toLowerCase();

// Split into words, keeping only letters (strip punctuation and non-letter characters)
// We'll use a regex to find sequences of letters
const wordRegex = /[a-z]+/g;
const matches = lowerContent.match(wordRegex);

if (!matches) {
  process.exit(0);
}

// Count frequency
const freq: Record<string, number> = {};
for (const word of matches) {
  freq[word] = (freq[word] || 0) + 1;
}

// Convert to array for sorting
const entries = Object.entries(freq);

// Sort by count descending, then alphabetically ascending for ties
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