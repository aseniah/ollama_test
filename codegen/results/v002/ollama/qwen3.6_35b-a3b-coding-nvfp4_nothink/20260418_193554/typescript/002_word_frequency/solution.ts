import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(inputPath, 'utf-8');

// Split into words, lowercase, strip punctuation
const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequency
const frequency: Record<string, number> = {};
for (const word of words) {
  frequency[word] = (frequency[word] || 0) + 1;
}

// Sort by count descending, then alphabetically ascending
const sortedWords = Object.entries(frequency).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}