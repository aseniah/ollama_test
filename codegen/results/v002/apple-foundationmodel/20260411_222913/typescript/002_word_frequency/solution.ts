import fs from 'fs';
import path from 'path';

// Function to count word frequency
function countWordFrequency(inputFile: string): Map<string, number> {
  const wordFrequency: Map<string, number> = new Map();

  // Read the file
  const content = fs.readFileSync(inputFile, 'utf8');

  // Split content into words
  const words = content.toLowerCase().split(/\W+/);

  // Count each word frequency
  words.forEach(word => {
    wordFrequency.set(word, (wordFrequency.get(word) || 0) + 1);
  });

  return wordFrequency;
}

// Main function
async function main() {
  const inputFile = 'input/text.txt';

  try {
    const wordFrequency = await countWordFrequency(inputFile);

    // Sort words first by count descending, then alphabetically ascending
    const sortedWords = wordFrequency.entries()
      .sort((a, b) => {
        if (a[1] !== b[1]) {
          return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
      })
      .map(([word, count]) => `${word}: ${count}`);

    // Print sorted word frequencies
    sortedWords.forEach(line => process.stdout.write(line + '\n'));
  } catch (error) {
    console.error('Error:', error);
  }
}

// Run the program
main();