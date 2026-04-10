import * as fs from 'node:fs';
import * as path from 'node:path';

/**
 * This script reads 'input/text.txt', counts word frequencies,
 * and outputs them according to the specified sorting rules.
 */

function main() {
  const filePath = path.join(process.cwd(), 'input', 'text.txt');

  let content: string;

  try {
    content = fs.readFileSync(filePath, 'utf-8');
  } catch (error) {
    // If file doesn't exist or can't be read, we exit silently 
    // as per "Do not output anything else" rule.
    process.exit(0);
  }

  // 1. Convert to lowercase
  // 2. Replace everything that isn't a letter (a-z) or whitespace with empty string
  // 3. Split by whitespace to get words
  const words = content
    .toLowerCase()
    .replace(/[^a-z\s]/g, '')
    .split(/\s+/);

  const frequencyMap: Map<string, number> = new Map();

  for (const word of words) {
    // Filter out empty strings resulting from multiple spaces or leading/trailing punctuation
    if (word.length > 0) {
      frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }
  }

  // Convert map to array of tuples [word, count] for sorting
  const sortedEntries = Array.from(frequencyMap.entries()).sort((a, b) => {
    const [wordA, countA] = a;
    const [wordB, countB] = b;

    // Sort by count descending
    if (countA !== countB) {
      return countB - countA;
    }

    // For ties, sort alphabetically ascending
    if (wordA < wordB) return -1;
    if (wordA > wordB) return 1;
    return 0;
  });

  // Output results
  for (const [word, count] of sortedEntries) {
    process.stdout.write(`${word}: ${count}\n`);
  }
}

main();