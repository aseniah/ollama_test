import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input/text.txt');
const text = fs.readFileSync(filePath, 'utf-8');

// Convert to lowercase, strip punctuation (keep only letters and whitespace), split into words
const words = text
  .toLowerCase()
  .replace(/[^a-z\s]/g, '')
  .split(/\s+/)
  .filter(word => word.length > 0);

// Count word frequencies
const wordCount = new Map<string, number>();
for (const word of words) {
  wordCount.set(word, (wordCount.get(word) || 0) + 1);
}

// Sort by count descending, then alphabetically ascending for ties
const sorted = Array.from(wordCount.entries()).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output results
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}