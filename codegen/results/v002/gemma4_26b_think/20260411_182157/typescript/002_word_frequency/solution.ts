import * as fs from 'fs';
import * as path from 'path';

/**
 * Task:
 * 1. Read the file `input/text.txt`.
 * 2. Convert all words to lowercase.
 * 3. Strip all punctuation (keep only letters).
 * 4. Count the frequency of each word.
 * 5. Output one line per unique word in the format: `word: count`.
 * 6. Sort by count descending; for ties, sort alphabetically ascending.
 * 7. Do not output anything else.
 */

function solve() {
  const filePath = path.join(process.cwd(), 'input', 'text.txt');

  let content: string;

  try {
    // Read file content synchronously
    content = fs.readFileSync(filePath, 'utf-8');
  } catch (error) {
    // If the file cannot be read, we exit silently as per requirements (do not output anything else)
    return;
  }

  // Convert content to lowercase
  const lowerContent = content.toLowerCase();

  // Use a regular expression to extract only sequences of letters (stripping punctuation)
  // This matches all substrings consisting of 'a' through 'z'
  const words = lowerContent.match(/[a-z]+/g);

  if (!words) {
    return;
  }

  // Map to store the frequency of each word
  const wordCounts = new Map<string, number>();

  for (const word of words) {
    const currentCount = wordCounts.get(word) || 0;
    wordCounts.set(word, currentCount + 1);
  }

  // Convert Map to an array of [word, count] pairs for sorting
  const entries = Array.from(wordCounts.entries());

  // Sort the entries:
  // 1. Primary Sort: Count descending (b[1] - a[1])
  // 2. Secondary Sort: Word ascending (a[0].localeCompare(b[0]))
  entries.sort((a, b) => {
    const countDiff = b[1] - a[1];
    if (countDiff !== 0) {
      return countDiff;
    }
    return a[0].localeCompare(b[0]);
  });

  // Output the results to stdout
  for (const [word, count] of entries) {
    process.stdout.write(`${word}: ${count}\n`);
  }
}

// Execute the solver
solve();