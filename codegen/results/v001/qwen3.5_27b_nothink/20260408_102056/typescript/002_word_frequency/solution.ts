import * as fs from 'fs';
import * as path from 'path';

function main() {
  const inputPath = path.join('input', 'text.txt');

  if (!fs.existsSync(inputPath)) {
    return;
  }

  const content = fs.readFileSync(inputPath, 'utf-8');

  // Normalize: lowercase, strip non-letters, split by whitespace
  const normalizedContent = content.toLowerCase().replace(/[^a-z\s]/g, '');
  const words = normalizedContent.split(/\s+/).filter(w => w.length > 0);

  const frequencyMap = new Map<string, number>();

  for (const word of words) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }

  const sortedEntries = Array.from(frequencyMap.entries())
    .sort((a, b) => {
      const countDiff = b[1] - a[1]; // Descending by count
      if (countDiff !== 0) return countDiff;
      return a[0].localeCompare(b[0]); // Ascending alphabetically
    });

  for (const [word, count] of sortedEntries) {
    console.log(`${word}: ${count}`);
  }
}

main();