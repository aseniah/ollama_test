import fs from 'fs';
import path from 'path';

// Read the file content
const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf8');

// Process the content
const words = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, ' ')
  .split(/\s+/)
  .filter(word => word.length > 0);

// Count frequencies
const wordCount: Record<string, number> = {};
for (const word of words) {
  wordCount[word] = (wordCount[word] || 0) + 1;
}

// Sort and output
const sortedEntries = Object.entries(wordCount)
  .sort(([wordA, countA], [wordB, countB]) => {
    if (countB !== countA) {
      return countB - countA; // Descending by count
    }
    return wordA.localeCompare(wordB); // Ascending by word
  });

for (const [word, count] of sortedEntries) {
  console.log(`${word}: ${count}`);
}