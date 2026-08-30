import * as fs from 'fs';
import * as path from 'path';

// Read the file
const filePath = path.join(process.cwd(), 'input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Split into words and process
const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequencies
const wordCounts = new Map<string, number>();
for (const word of words) {
  wordCounts.set(word, (wordCounts.get(word) || 0) + 1);
}

// Convert to array and sort
const sortedWords = Array.from(wordCounts.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Sort by count descending
    }
    return a[0].localeCompare(b[0]); // Sort alphabetically ascending for ties
  });

// Output
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}