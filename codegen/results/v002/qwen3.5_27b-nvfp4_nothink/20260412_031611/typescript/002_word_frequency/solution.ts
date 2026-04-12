import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'text.txt');

try {
  const data = fs.readFileSync(filePath, 'utf8');
  
  // Split by non-alphabetic characters to get raw tokens
  const rawWords = data.toLowerCase().split(/[^a-z]+/);
  
  const wordCounts = new Map<string, number>();

  for (const word of rawWords) {
    if (word.length > 0) {
      wordCounts.set(word, (wordCounts.get(word) || 0) + 1);
    }
  }

  // Convert map to array for sorting
  const sortedEntries = Array.from(wordCounts.entries());

  // Sort: count descending, then word ascending
  sortedEntries.sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Descending by count
    }
    return a[0].localeCompare(b[0]); // Ascending alphabetically
  });

  // Output results
  for (const [word, count] of sortedEntries) {
    console.log(`${word}: ${count}`);
  }
} catch (error) {
  // Do not output anything else, even errors, per requirements
}