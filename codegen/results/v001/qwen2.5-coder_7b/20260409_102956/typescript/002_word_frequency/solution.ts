import * as fs from 'fs';
import * as path from 'path';

// Function to remove punctuation and convert text to lowercase
function cleanText(text: string): string {
  return text.replace(/[^\w\s]/g, '').toLowerCase();
}

// Function to count word frequencies
async function countWordFrequencies(filePath: string): Promise<void> {
  try {
    const fileContent = await fs.promises.readFile(filePath, 'utf8');
    const cleanTextStr = cleanText(fileContent);
    const words = cleanTextStr.split(/\s+/).filter(word => word.length > 0);

    const frequencyMap: { [word: string]: number } = {};

    for (const word of words) {
      if (frequencyMap[word]) {
        frequencyMap[word]++;
      } else {
        frequencyMap[word] = 1;
      }
    }

    // Sort the words by frequency descending, then alphabetically ascending
    const sortedWords = Object.keys(frequencyMap)
      .sort((a, b) => {
        if (frequencyMap[a] === frequencyMap[b]) {
          return a.localeCompare(b);
        }
        return frequencyMap[b] - frequencyMap[a];
      });

    // Output each word with its count
    for (const word of sortedWords) {
      console.log(`${word}: ${frequencyMap[word]}`);
    }

  } catch (err) {
    console.error(`Error reading file: ${err.message}`);
  }
}

// Main function to run the script
async function main(): Promise<void> {
  const filePath = path.join(__dirname, 'input', 'text.txt');
  await countWordFrequencies(filePath);
}

main();