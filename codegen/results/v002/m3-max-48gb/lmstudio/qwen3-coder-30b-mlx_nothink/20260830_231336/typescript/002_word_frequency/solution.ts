import fs from 'fs';
import path from 'path';

// Read the input file
const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf8');

// Process the content
const words = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, ' ') // Replace all non-letter characters with spaces
  .split(/\s+/) // Split by whitespace
  .filter(word => word.length > 0); // Remove empty strings

// Count word frequencies
const wordCount: { [key: string]: number } = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Sort by count descending, then alphabetically ascending
const sortedEntries = Object.entries(wordCount).sort(([wordA, countA], [wordB, countB]) => {
  if (countB !== countA) {
    return countB - countA; // Sort by count descending
  }
  return wordA.localeCompare(wordB); // Sort alphabetically ascending
});

// Output results
for (const [word, count] of sortedEntries) {
  console.log(`${word}: ${count}`);
}