import fs from 'fs';
import path from 'path';

// Function to count word frequencies
function countWordFrequencies(inputFilePath: string): string {
  const wordCount: Record<string, number> = {};

  // Read the file contents
  fs.readFile(inputFilePath, 'utf8', (err, data) => {
    if (err) {
      console.error('Error reading file:', err);
      return;
    }

    // Split the data into words
    const words = data.toLowerCase().split(/\W+/);

    // Count each word
    words.forEach(word => {
      if (wordCount[word]) {
        wordCount[word]++;
      } else {
        wordCount[word] = 1;
      }
    });

    // Sort words first by count descending, then alphabetically ascending
    const sortedWords = Object.entries(wordCount).sort((a, b) => {
      const countDiff = b[1] - a[1];
      if (countDiff === 0) {
        return a[0].localeCompare(b[0]);
      }
      return countDiff;
    });

    // Prepare the output string
    const output = sortedWords.map(entry => `${entry[0]}: ${entry[1]}`).join('\n');

    return output;
  });
}

// Main execution
if (process.argv.length !== 2) {
  console.error('Usage: node script.tsx <input file path>');
} else {
  const inputFilePath = process.argv[2];
  if (!path.existsSync(inputFilePath)) {
    console.error('Input file not found:', inputFilePath);
  } else {
    console.log(countWordFrequencies(inputFilePath));
  }
}