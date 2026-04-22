import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Split by whitespace and punctuation, keep only letters
const words = content
  .toLowerCase()
  .split(/[^a-zA-Z]+/)
  .filter(word => word.length > 0);

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

for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}