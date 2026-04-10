import * as fs from 'fs';
import * as path from 'path';

const filePath = process.argv[2] || 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Convert to lowercase and keep only letters (and whitespace for splitting)
const cleaned = content.toLowerCase().replace(/[^a-z\s]/g, '');
const words = cleaned.split(/\s+/).filter(w => w.length > 0);

// Count frequency
const wordCount: { [word: string]: number } = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Convert to array and sort by count descending, then alphabetically ascending for ties
const entries = Object.entries(wordCount);
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1]; // Count descending
  }
  return a[0].localeCompare(b[0]); // Alphabetically ascending
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}