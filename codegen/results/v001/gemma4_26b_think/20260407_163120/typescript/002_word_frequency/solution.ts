import * as fs from 'fs';
import * as path from 'path';

/**
 * The script reads 'input/text.txt', processes the text to count word frequencies,
 * and outputs the results sorted by frequency (descending) and then alphabetically (ascending).
 */

function main() {
  const inputPath = path.join('input', 'text.txt');

  try {
    // Check if file exists before attempting to read
    if (!fs.existsSync(inputPath)) {
      return;
    }

    const content = fs.readFileSync(inputPath, 'utf-8');

    // 1. Convert to lowercase
    // 2. Replace all non-letter characters with spaces to act as delimiters
    // 3. This effectively strips punctuation and keeps only letters
    const cleanContent = content
      .toLowerCase()
      .replace(/[^a-z]/g, ' ');

    // Split by whitespace and filter out empty strings
    const words = cleanContent.split(/\s+/).filter(word => word.length > 0);

    const wordCounts = new Map<string, number>();

    // Count frequencies
    for (const word of words) {
      const count = wordCounts.get(word) || 0;
      wordCounts.set(word, count + 1);
    }

    // Convert Map to an array of [word, count] tuples for sorting
    const sortedEntries = Array.from(wordCounts.entries());

    // Sort: 
    // Primary: count descending (b[1] - a[1])
    // Secondary: word ascending (a[0].localeCompare(b[0]))
    sortedEntries.sort((a, b) => {
      if (b[1] !== a[1]) {
        return b[1] - a[1];
      }
      return a[0].localeCompare(b[0]);
    });

    // Output the result in the format "word: count"
    for (const [word, count] of sortedEntries) {
      process.stdout.write(`${word}: ${count}\n`);
    }
  } catch (error) {
    // Per instructions, do not output anything else (including errors)
    // unless it's the required output.
  }
}

main();