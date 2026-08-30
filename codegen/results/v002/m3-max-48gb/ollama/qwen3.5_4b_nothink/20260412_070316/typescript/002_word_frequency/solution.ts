import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.resolve(process.argv[2] || 'input/text.txt');

// Read file content
const content = fs.readFileSync(inputFile, 'utf-8');

// Process the content:
// 1. Convert to lowercase
// 2. Remove all non-letter characters (keep only a-z)
// 3. Split into words
// 4. Filter out empty strings

let processedText = content.toLowerCase();

// Replace everything that is NOT a letter with nothing
processedText = processedText.replace(/[^a-z]/g, '');

// Split into words
const words = processedText.split(' ');

// Count frequencies
const wordCounts: Map<string, number> = new Map();

for (const word of words) {
  if (word.length > 0) {
    const count = wordCounts.get(word) || 0;
    wordCounts.set(word, count + 1);
  }
}

// Convert to array for sorting
const sortedWords = Array.from(wordCounts.entries())
  .sort((a, b) => {
    // Sort by count descending
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    // Sort alphabetically ascending for ties
    return a[0].localeCompare(b[0]);
  })
  .map(([word, count]) => `${word}: ${count}`);

// Output result
console.log(sortedWords.join('\n'));