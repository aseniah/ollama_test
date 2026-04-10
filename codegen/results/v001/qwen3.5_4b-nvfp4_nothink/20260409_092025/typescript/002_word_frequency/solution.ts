import { readFileSync } from 'fs';
import { basename } from 'path';

const filePath = process.argv[2];

if (!filePath) {
  console.error('Usage: node script.js <filename>');
  process.exit(1);
}

const content = readFileSync(filePath, 'utf-8');
const lines = content.split('\n');

const wordCounts = new Map<string, number>();

for (const line of lines) {
  let text = line.toLowerCase();
  
  // Remove all non-letter characters
  // Replace anything that is not a letter with nothing (empty string)
  text = text.replace(/[^a-z]/g, '');
  
  const words = text.split(/\s+/);

  for (const word of words) {
    if (word.length > 0) {
      const count = wordCounts.get(word) || 0;
      wordCounts.set(word, count + 1);
    }
  }
}

// Sort entries: first by count descending, then by word ascending
const sortedEntries = Array.from(wordCounts.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Descending count
    }
    return a[0].localeCompare(b[0]); // Ascending word
  });

const outputLines = sortedEntries.map(([word, count]) => `${word}: ${count}`);
console.log(outputLines.join('\n'));