import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(inputPath, 'utf-8');

// Normalize: lowercase and split into words
// Keep only letters (a-z), discard everything else
const text = content.toLowerCase();
const words = text.split(/\s+/).map(w => w.replace(/[^a-z]/g, '')).filter(w => w.length > 0);

// Count frequencies
const counts = new Map<string, number>();
for (const word of words) {
  counts.set(word, (counts.get(word) || 0) + 1);
}

// Sort: by count descending, then alphabetically ascending for ties
const sorted = Array.from(counts.entries()).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}