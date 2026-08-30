import fs from 'fs';
import path from 'path';

const inputPath = process.argv[2];

if (!inputPath) {
  console.log('Usage: ts-node <input/text.txt>');
  process.exit(1);
}

const content = fs.readFileSync(inputPath, 'utf-8');

// Convert to lowercase
const lowerContent = content.toLowerCase();

// Remove all characters that are not letters (a-z)
const wordsOnly = lowerContent.replace(/[^a-z]/g, '');

// Split into words
const words = wordsOnly.split(/\s+/);

// Count frequencies
const counts: Record<string, number> = {};
for (const word of words) {
  if (word.length > 0) {
    counts[word] = (counts[word] || 0) + 1;
  }
}

// Convert to array and sort:
// 1. By count descending
// 2. By word ascending (alphabetical) for ties
const sortedWords = Object.entries(counts).sort((a, b) => {
  const countA = parseInt(a[1], 10);
  const countB = parseInt(b[1], 10);
  if (countB !== countA) {
    return countB - countA; // Descending count
  } else {
    return a[0].localeCompare(b[0]); // Ascending alphabetical
  }
});

// Output results
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}