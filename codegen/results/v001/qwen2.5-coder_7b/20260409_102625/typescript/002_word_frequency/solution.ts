const fs = require('fs');
const path = require('path');

// Read the input file
const filePath = path.join(__dirname, 'input', 'text.txt');
const fileContent = fs.readFileSync(filePath, 'utf-8');

// Convert all words to lowercase and strip punctuation
const cleanedContent = fileContent.toLowerCase().replace(/[^\w\s]/g, '');

// Split content into words
const words = cleanedContent.split(/\s+/);

// Create a map to count word frequencies
const wordCountMap: { [key: string]: number } = {};

words.forEach(word => {
    if (word) {
        if (wordCountMap[word]) {
            wordCountMap[word]++;
        } else {
            wordCountMap[word] = 1;
        }
    }
});

// Sort the words by frequency and then alphabetically
const sortedWords = Object.entries(wordCountMap)
    .sort((a, b) => {
        if (a[1] === b[1]) {
            return a[0].localeCompare(b[0]);
        }
        return b[1] - a[1];
    })
    .map(([word, count]) => `${word}: ${count}`);

// Output the results
sortedWords.forEach(line => console.log(line));