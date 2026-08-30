import * as fs from 'fs';

// Read the file
const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Process words - convert to lowercase and extract only letters
const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequencies
const frequencyMap: Map<string, number> = new Map();
for (const word of words) {
  frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
}

// Convert to array and sort
const sortedEntries = Array.from(frequencyMap.entries())
  .sort((a, b) => {
    // Sort by count descending
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    // For ties, sort alphabetically ascending
    return a[0].localeCompare(b[0]);
  });

// Output
for (const [word, count] of sortedEntries) {
  console.log(`${word}: ${count}`);
}