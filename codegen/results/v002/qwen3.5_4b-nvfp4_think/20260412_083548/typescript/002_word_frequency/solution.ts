import fs from 'fs';
import path from 'path';

const filePath = 'input/text.txt';

// Read file content
const content = fs.readFileSync(filePath, 'utf-8');

// Split into potential words and filter using regex (keep only letters)
const words = content.match(/[a-z]+/gi);

if (!words) {
  process.exit(0);
}

// Count word frequencies using a Map for efficiency
const counts: Map<string, number> = new Map<string, number>();

for (const word of words!) {
  const cleanWord = word.toLowerCase(); // Already in lowercase due to regex flags
  
  // Increment count
  const existingCount = counts.get(cleanWord) || 0;
  counts.set(cleanWord, existingCount + 1);
}

// Convert Map to array and sort properly
const wordArray = Array.from(counts.entries());

// Sort by count descending, then alphabetically ascending
wordArray.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1]; // Descending by count
  return a[0].localeCompare(b[0]);          // Alphabetically ascending for ties
});

// Output results in the required format
for (const [word, count] of wordArray) {
  console.log(`${word}: ${count}`);
}