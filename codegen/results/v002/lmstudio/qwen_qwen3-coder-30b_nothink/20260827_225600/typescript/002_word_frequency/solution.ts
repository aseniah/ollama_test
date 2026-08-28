import fs from 'fs';
import path from 'path';

// Read the file content
const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Process the content to count word frequencies
const wordCounts: Record<string, number> = {};

// Split by whitespace and process each word
const words = content.split(/\s+/);

for (const word of words) {
  // Remove punctuation and convert to lowercase
  const cleanWord = word.replace(/[^a-zA-Z]/g, '').toLowerCase();
  
  // Skip empty strings
  if (cleanWord) {
    wordCounts[cleanWord] = (wordCounts[cleanWord] || 0) + 1;
  }
}

// Sort by count descending, then alphabetically ascending
const sortedEntries = Object.entries(wordCounts).sort(([wordA, countA], [wordB, countB]) => {
  if (countB !== countA) {
    return countB - countA; // Sort by count descending
  }
  return wordA.localeCompare(wordB); // Sort alphabetically ascending
});

// Output the results
for (const [word, count] of sortedEntries) {
  console.log(`${word}: ${count}`);
}