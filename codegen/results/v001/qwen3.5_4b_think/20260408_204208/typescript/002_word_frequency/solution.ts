import * as fs from 'fs';
import * as path from 'path';

// Read the file
const filePath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Process the content: lowercase and keep only letters
const words = content
  .toLowerCase()
  .split(/[^a-z\s]+/)
  .filter(word => word.length > 0);

// Count frequency
const frequency: Map<string, number> = new Map();
for (const word of words) {
  frequency.set(word, (frequency.get(word) || 0) + 1);
}

// Convert to array and sort
const sorted = Array.from(frequency.entries())
  .sort((a, b) => {
    // Sort by count descending
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    // Sort alphabetically ascending for ties
    return a[0].localeCompare(b[0]);
  });

// Output
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}