const fs = require('fs');
const path = require('path');

// Read the input file
const inputPath = process.argv[2] || 'input/text.txt';
let content = '';

try {
  content = fs.readFileSync(inputPath, 'utf8');
} catch (error) {
  console.error(`Error reading file: ${error.message}`);
  process.exit(1);
}

// Helper function to remove non-letter characters
function normalizeWord(word) {
  return word.replace(/[^a-zA-Z]/g, '');
}

// Split content into words, normalize, and count
const words = content
  .split(/\s+/) // Split by whitespace
  .map(normalizeWord) // Normalize each word
  .filter(word => word.length > 0); // Remove empty strings

const wordCountMap = {};

for (const word of words) {
  if (wordCountMap[word]) {
    wordCountMap[word]++;
  } else {
    wordCountMap[word] = 1;
  }
}

// Convert map to array and sort
const sortedWords = Object.entries(wordCountMap)
  .sort((a, b) => {
    // Sort by count descending
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    // Sort by word alphabetically ascending (already lowercase due to normalization)
    return a[0].localeCompare(b[0]);
  })
  .map(entry => `${entry[0]}: ${entry[1]}`);

// Output the result
console.log(sortedWords.join('\n'));