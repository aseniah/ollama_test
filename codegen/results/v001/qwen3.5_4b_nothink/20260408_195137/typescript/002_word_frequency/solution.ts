const fs = require('fs');
const path = require('path');

const inputPath = process.argv[2];
const fileContent = fs.readFileSync(inputPath, 'utf8');

// Split into lines, then words
const lines = fileContent.split('\n');
const allWords = lines.map(line => {
  const words = line.match(/[a-z]+/gi);
  return words ? words.join(' ') : '';
}).join(' ');

// Extract all words
const wordList = allWords.match(/[a-z]+/g);
const counts = {};

if (!wordList) {
  return;
}

wordList.forEach(word => {
  counts[word] = (counts[word] || 0) + 1;
});

// Convert to array and sort
const wordCounts = Object.entries(counts).map(([word, count]) => ({ word, count }));

// Sort by count descending, then word ascending
wordCounts.sort((a, b) => {
  if (b.count !== a.count) {
    return b.count - a.count;
  }
  return a.word < b.word ? -1 : 1;
});

wordCounts.forEach(({ word, count }) => {
  console.log(`${word}: ${count}`);
});
```