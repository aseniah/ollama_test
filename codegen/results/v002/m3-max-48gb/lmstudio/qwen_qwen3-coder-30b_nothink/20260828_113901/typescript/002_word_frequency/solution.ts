import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'text.txt');

// Read the file content
const content = fs.readFileSync(filePath, 'utf-8');

// Split into words and process them
const words = content
  .split(/\s+/)
  .map(word => 
    // Remove punctuation and convert to lowercase
    word.toLowerCase().replace(/[^a-z]/g, '')
  )
  .filter(word => word.length > 0); // Filter out empty strings

// Count word frequencies
const wordCount: { [key: string]: number } = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Sort by count descending, then alphabetically ascending
const sortedEntries = Object.entries(wordCount).sort(([wordA, countA], [wordB, countB]) => {
  if (countB !== countA) {
    return countB - countA; // Descending count
  }
  return wordA.localeCompare(wordB); // Ascending alphabetical
});

// Output results
for (const [word, count] of sortedEntries) {
  console.log(`${word}: ${count}`);
}