import { readFileSync } from 'fs';
import { join } from 'path';

// Read the file
const filePath = join('input', 'text.txt');
const content = readFileSync(filePath, 'utf-8');

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

// Sort by count (descending) then by word (ascending)
const sortedWords = Object.entries(wordCount)
  .sort(([wordA, countA], [wordB, countB]) => {
    if (countB !== countA) {
      return countB - countA; // Sort by count descending
    }
    return wordA.localeCompare(wordB); // Sort by word ascending
  });

// Output results
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}