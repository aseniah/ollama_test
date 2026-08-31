import fs from 'fs';
import path from 'path';

// Read the file
const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Process the content
const words = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, ' ')
  .split(/\s+/)
  .filter(word => word.length > 0);

// Count frequencies
const wordCount: { [key: string]: number } = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Sort and output
const sortedWords = Object.entries(wordCount)
  .sort(([wordA, countA], [wordB, countB]) => {
    if (countB !== countA) {
      return countB - countA; // Sort by count descending
    }
    return wordA.localeCompare(wordB); // Sort alphabetically ascending for ties
  });

for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}