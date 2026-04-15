import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(inputPath, 'utf-8');

// Normalize: convert to lowercase
const lowerContent = content.toLowerCase();

// Extract words: keep only letters, split by non-letters
const words = lowerContent.match(/[a-z]+/g) || [];

// Count frequencies
const freqMap = new Map<string, number>();
for (const word of words) {
  freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

// Convert to array for sorting
const entries = Array.from(freqMap.entries());

// Sort by count descending, then by word ascending
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1]; // count descending
  }
  return a[0].localeCompare(b[0]); // word ascending
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}