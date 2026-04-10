import * as fs from 'fs';
import * as path from 'path';

// Read file content
const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Split into words (by whitespace and punctuation)
const words = content.match(/\b[a-zA-Z]+\b/g) || [];

// Count frequency (lowercase only)
const wordCount: Record<string, number> = {};
words.forEach(word => {
  const lowerWord = word.toLowerCase();
  wordCount[lowerWord] = (wordCount[lowerWord] || 0) + 1;
});

// Convert to array for sorting
const sortedWords = Object.entries(wordCount)
  .sort((a, b) => {
    const countA = parseInt(a[1]);
    const countB = parseInt(b[1]);
    if (countB !== countA) {
      return countB - countA; // Descending by count
    }
    return a[0].localeCompare(b[0]); // Ascending alphabetically
  });

// Output results
sortedWords.forEach(([word, count]) => {
  console.log(`${word}: ${count}`);
});