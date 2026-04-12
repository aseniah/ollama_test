import * as fs from 'fs';
import * as path from 'path';

/**
 * The task is to read a specific file 'input/text.txt', 
 * count the frequency of each word, and output the results
 * sorted by frequency (descending) and then alphabetically (ascending).
 * 
 * Rules:
 * - Convert all words to lowercase.
 * - Strip all punctuation (keep only letters).
 * - Output format: 'word: count'
 * - Sort: count descending, then alphabetical ascending.
 */

function solve() {
  // The problem specifies the file path 'input/text.txt'
  const filePath = path.join('input', 'text.txt');
  let content: string;

  try {
    // Read file content using synchronous reading
    content = fs.readFileSync(filePath, 'utf8');
  } catch (err) {
    // If the file cannot be read (e.g., not found), we exit silently
    // as there is no output required for error states in the prompt.
    return;
  }

  // Split the content into tokens based on whitespace (spaces, tabs, newlines)
  const tokens = content.split(/\s+/);
  const wordCounts = new Map<string, number>();

  for (const token of tokens) {
    // Rule 1: Convert to lowercase
    // Rule 2: Strip all punctuation (keep only letters a-z)
    const cleanWord = token.toLowerCase().replace(/[^a-z]/g, '');

    // Only process if the word is not empty after stripping punctuation
    if (cleanWord.length > 0) {
      const currentCount = wordCounts.get(cleanWord) || 0;
      wordCounts.set(cleanWord, currentCount + 1);
    }
  }

  // Convert the Map into an array of [word, count] pairs for sorting
  const sortedEntries = Array.from(wordCounts.entries()).sort((a, b) => {
    const [wordA, countA] = a;
    const [wordB, countB] = b;

    // Primary Sort: Frequency descending
    if (countB !== countA) {
      return countB - countA;
    }

    // Secondary Sort: Alphabetical ascending (using localeCompare for string comparison)
    return wordA.localeCompare(wordB);
  });

  // Output the results to stdout
  for (const [word, count] of sortedEntries) {
    process.stdout.write(`${word}: ${count}\n`);
  }
}

// Execute the solver
solve();