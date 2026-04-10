import * as fs from 'fs';
import * as path from 'path';

/**
 * Word Frequency Counter
 * 
 * This script reads a text file, processes the words based on specific
 * transformation rules, and outputs the frequency of each word.
 * 
 * Rules:
 * - Convert all words to lowercase
 * - Strip all punctuation (keep only letters)
 * - Output format: "word: count"
 * - Sort: Frequency (descending), then Alphabetical (ascending)
 */

function main() {
  // The problem specifies reading 'input/text.txt'. 
  // We allow an optional command-line argument for flexibility.
  const targetPath = process.argv[2] || path.join('input', 'text.txt');

  try {
    // Ensure the file exists before attempting to read
    if (!fs.existsSync(targetPath)) {
      return;
    }

    const content = fs.readFileSync(targetPath, 'utf-8');

    // 1. Convert all content to lowercase
    // 2. Use a regular expression to replace everything that is NOT a letter [a-z] 
    //    with a space. This effectively strips punctuation and digits.
    // 3. Split the string into an array using whitespace as the delimiter.
    // 4. Filter out any empty strings resulting from multiple delimiters.
    const words = content
      .toLowerCase()
      .replace(/[^a-z]/g, ' ')
      .split(/\s+/)
      .filter((word) => word.length > 0);

    const frequencyMap = new Map<string, number>();

    // Populate the frequency map
    for (const word of words) {
      const currentCount = frequencyMap.get(word) || 0;
      frequencyMap.set(word, currentCount + 1);
    }

    // Convert the map to an array of tuples [word, count] for sorting
    const sortedEntries = Array.from(frequencyMap.entries()).sort((a, b) => {
      const [wordA, countA] = a;
      const [wordB, countB] = b;

      // Primary Sort: Count Descending
      if (countB !== countA) {
        return countB - countA;
      }

      // Secondary Sort: Alphabetical Ascending
      return wordA.localeCompare(wordB);
    });

    // Final Output: One line per unique word in the format 'word: count'
    for (const [word, count] of sortedEntries) {
      process.stdout.write(`${word}: ${count}\n`);
    }
  } catch (error) {
    // We catch errors but do not output anything to stdout to comply with 
    // the "Do not output anything else" rule.
  }
}

// Run the application
main();