import * as fs from 'fs';
import * as path from 'path';

function main() {
  const filePath = path.resolve(__dirname, 'input', 'text.txt');
  const content = fs.readFileSync(filePath, 'utf-8');

  // Split into words: split on whitespace, then strip punctuation from each
  const words = content.split(/\s+/).filter(w => w.length > 0);

  const frequency: Record<string, number> = {};

  for (const word of words) {
    // Strip all non-letter characters, keep only [a-zA-Z]
    const cleaned = word.replace(/[^a-zA-Z]/g, '');
    if (cleaned.length === 0) continue;

    const lower = cleaned.toLowerCase();
    frequency[lower] = (frequency[lower] || 0) + 1;
  }

  // Sort by count descending, then alphabetically ascending for ties
  const sorted = Object.entries(frequency).sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
  });

  // Output
  for (const [word, count] of sorted) {
    console.log(`${word}: ${count}`);
  }
}

main();