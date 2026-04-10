import fs from 'fs';
import path from 'path';

const filePath = path.resolve(__dirname, 'input/text.txt');

// Read the file content
const content = fs.readFileSync(filePath, 'utf8');

// Split content into words and normalize them
const words = content.toLowerCase().split(/\s+/).map(word => {
  const regex = /[^a-z]/gi;
  return word.replace(regex, '');
});

// Count frequencies
const wordCount: Record<string, number> = new Map();
words.forEach(word => wordCount.set(word, (wordCount.get(word) || 0) + 1));

// Sort words first by count descending, then alphabetically ascending
const sortedWords = [...wordCount.keys()].sort((a, b) => {
  const countA = wordCount.get(a);
  const countB = wordCount.get(b);
  if (countA !== countB) {
    return countA - countB;
  }
  return a.localeCompare(b);
});

// Output the result
sortedWords.forEach(word => console.log(`${word}: ${wordCount.get(word)}`));