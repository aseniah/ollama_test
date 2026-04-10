import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads 'input/text.txt', counts word frequencies,
 * and outputs the results sorted by count (desc) and then alphabetically (asc).
 */

function main() {
  const filePath = path.join(process.cwd(), 'input', 'text.txt');

  // Check if file exists
  if (!fs.existsSync(filePath)) {
    process.exit(0);
  }

  try {
    const content = fs.readFileSync(filePath, 'utf-8');

    // Regex explanation:
    // We replace anything that is NOT a letter or whitespace with an empty string.
    // This strips punctuation while preserving spaces.
    const cleanContent = content.toLowerCase().replace(/[^a-z\s]/g, ' ');

    // Split by whitespace and filter out empty strings
    const words = cleanContent.split(/\s+/).filter(word => word.length > 0);

    const frequencyMap: Map<string, number> = new Map();

    for (const word of words) {
      const count = frequencyMap.get(word) || 0;
      frequencyMap.set(word, count + 1);
    }

    // Convert map to array for sorting
    const sortedEntries = Array.from(frequencyMap.entries()).sort((a, b) => {
      const [wordA, countA] = a;
      const [wordB, countB] = b;

      // Primary sort: Count descending
      if (countB !== countA) {
        return countB - countA;
      }

      // Secondary sort: Alphabetical ascending
      return wordA.localeCompare(wordB);
    });

    // Output the results
    for (const [word, count] of sortedEntries) {
      process.stdout.write(`${word}: ${count}\n`);
    }
  } catch (error) {
    // Silently exit or handle error as per requirements (no extra output)
    process.exit(1);
  }
}

main();