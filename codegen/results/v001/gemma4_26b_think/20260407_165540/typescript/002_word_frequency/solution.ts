import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads 'input/text.txt', processes the text to count word frequencies,
 * and outputs the results according to specific sorting and formatting rules.
 */
function main() {
  const filePath = path.join(process.cwd(), 'input', 'text.txt');

  try {
    // Check if the file exists before attempting to read
    if (!fs.existsSync(filePath)) {
      return;
    }

    const content = fs.readFileSync(filePath, 'utf8');

    // 1. Convert to lowercase
    // 2. Replace anything that is not a letter (a-z) with a space
    //    (This strips punctuation and numbers as requested)
    // 3. Split by whitespace to get individual words
    // 4. Filter out empty strings resulting from multiple spaces/punctuation
    const words = content
      .toLowerCase()
      .replace(/[^a-z]/g, ' ')
      .split(/\s+/)
      .filter((word) => word.length > 0);

    // Count the frequency of each word using a Map
    const wordCounts = new Map<string, number>();
    for (const word of words) {
      const currentCount = wordCounts.get(word) || 0;
      wordCounts.set(word, currentCount + 1);
    }

    // Convert the Map entries to an array for sorting
    const sortedEntries = Array.from(wordCounts.entries());

    // Sorting logic:
    // Primary: Count descending (b[1] - a[1])
    // Secondary: Alphabetical ascending (a[0].localeCompare(b[0]))
    sortedEntries.sort((a, b) => {
      const countA = a[1];
      const countB = b[1];
      const wordA = a[0];
      const wordB = b[0];

      if (countB !== countA) {
        return countB - countA;
      }
      return wordA.localeCompare(wordB);
    });

    // Output the results in the format: word: count
    for (const [word, count] of sortedEntries) {
      process.stdout.write(`${word}: ${count}\n`);
    }
  } catch (error) {
    // If an error occurs (e.g., permission issues), we exit silently 
    // to adhere to the "Do not output anything else" rule.
    process.exit(1);
  }
}

main();