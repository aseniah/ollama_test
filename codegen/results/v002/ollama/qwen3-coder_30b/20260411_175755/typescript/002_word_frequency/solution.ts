import { readFileSync } from 'fs';
import { join } from 'path';

// Read the file content
const filePath = join('input', 'text.txt');
const content = readFileSync(filePath, 'utf-8');

// Convert to lowercase and extract words (only letters)
const words = content
  .toLowerCase()
  .match(/[a-z]+/g) || [];

// Count word frequencies
const wordCount: { [key: string]: number } = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Sort by count (descending) then alphabetically (ascending)
const sortedWords = Object.entries(wordCount)
  .sort(([wordA, countA], [wordB, countB]) => {
    if (countB !== countA) {
      return countB - countA; // Sort by count descending
    }
    return wordA.localeCompare(wordB); // Sort alphabetically ascending
  });

// Output results
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}