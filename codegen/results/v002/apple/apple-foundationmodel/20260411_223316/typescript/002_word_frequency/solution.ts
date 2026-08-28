import fs from 'fs';
import path from 'path';

// Function to count word frequencies
const countWordFrequencies = (inputFilePath: string): number => {
  const wordCount: { [key: string]: number } = {};

  // Read the file content
  fs.readFile(inputFilePath, 'utf8', (err, data) => {
    if (err) {
      throw new Error(`Error reading file: ${err.message}`);
    }

    // Split the content into words
    const words = data.toLowerCase().split(/\W+/).filter(word => word);

    // Count the frequency of each word
    words.forEach(word => {
      if (wordCount[word]) {
        wordCount[word]++;
      } else {
        wordCount[word] = 1;
      }
    });

    // Return the total word count
    return Object.keys(wordCount).length;
  });
};

// Function to format and print the results
const printWordFrequencies = (wordCount: number): void => {
  // Sort by count descending, then alphabetically ascending
  const sortedWords = Object.entries(wordCount)
    .sort((a, b) => {
      const countA = a[1];
      const countB = b[1];
      if (countA !== countB) {
        return countA - countB;
      } else {
        return a[0].localeCompare(b[0]);
      }
    })
    .map(entry => `${entry[0]}: ${entry[1]}`);

  // Print each word with its count
  console.log(...sortedWords);
};

// Main function to execute the process
const main = () => {
  const inputFilePath = path.join(__dirname, 'input/text.txt');

  try {
    const totalWords = countWordFrequencies(inputFilePath);
    printWordFrequencies(totalWords);
  } catch (error) {
    console.error(error.message);
  }
};

// Execute the main function
main();