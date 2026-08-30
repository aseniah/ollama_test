import * as fs from 'fs';

// Read the file
const content = fs.readFileSync('input/text.txt', 'utf8');

// Split into words and process
const words = content.toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

// Count frequencies
const frequency: Record<string, number> = {};
for (const word of words) {
  frequency[word] = (frequency[word] || 0) + 1;
}

// Sort by count descending, then alphabetically ascending
const sorted = Object.entries(frequency).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1]; // count descending
  }
  return a[0].localeCompare(b[0]); // alphabetically ascending
});

// Output
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}