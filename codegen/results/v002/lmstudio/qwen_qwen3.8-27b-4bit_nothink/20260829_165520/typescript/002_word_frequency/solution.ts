import * as fs from 'fs';

const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Tokenize: split on non-letter characters, then convert to lowercase
const words = content
  .toLowerCase()
  .split(/[^a-z]+/)
  .filter(word => word.length > 0);

// Count frequency
const freq: Map<string, number> = new Map();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array for sorting
const entries: [string, number][] = Array.from(freq.entries());

// Sort by count descending, then alphabetically ascending
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
process.stdout.write(output + '\n');