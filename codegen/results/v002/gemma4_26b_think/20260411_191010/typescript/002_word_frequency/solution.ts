import * as fs from 'fs';

/**
 * This program reads a specific file 'input/text.txt', processes the text 
 * to count the frequency of each word (after cleaning), and outputs 
 * the results sorted by frequency (descending) and then alphabetically (ascending).
 */

function solve(): void {
  const filePath = 'input/text.txt';

  try {
    // Check if the file exists before attempting to read
    if (!fs.existsSync(filePath)) {
      return;
    }

    // Read the file content
    const content = fs.readFileSync(filePath, 'utf-8');

    // Split text into tokens by any whitespace character
    const tokens = content.split(/\s+/);
    const wordCounts = new Map<string, number>();

    for (const token of tokens) {
      // 1. Convert all words to lowercase
      // 2. Strip all punctuation (keep only letters a-z)
      const cleanedWord = token.toLowerCase().replace(/[^a-z]/g, '');

      // Only process if the resulting string is not empty
      if (cleanedWord.length > 0) {
        const currentCount = wordCounts.get(cleanedWord) || 0;
        wordCounts.set(cleanedWord, currentCount + 1);
      }
    }

    // Convert Map entries to an array for sorting
    const sortedEntries = Array.from(wordCounts.entries());

    // Sort by count descending; for ties, sort alphabetically ascending
    sortedEntries.sort((a, b) => {
      const countA = a[1];
      const countB = b[1];
      const wordA = a[0];
      const wordB = b[0];

      if (countB !== countA) {
        return countB - countA; // Descending count
      }
      // Alphabetical ascending for ties
      return wordA.localeCompare(wordB);
    });

    // Output the results in the format 'word: count'
    for (const [word, count] of sortedEntries) {
      process.stdout.write(`${word}: ${count}\n`);
    }
  } catch (err) {
    // In case of unexpected errors, do not output anything else per instructions.
  }
}

// Execute the solver
solve();