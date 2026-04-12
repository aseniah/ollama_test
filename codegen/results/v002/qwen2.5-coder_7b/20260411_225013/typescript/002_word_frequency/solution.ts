import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, 'input', 'text.txt');

fs.readFile(inputFilePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading the file:', err);
    return;
  }

  const words = data.toLowerCase().match(/[a-z]+/g) || [];
  const wordCount: { [key: string]: number } = {};

  for (const word of words) {
    wordCount[word] = (wordCount[word] || 0) + 1;
  }

  const sortedWords = Object.entries(wordCount)
    .sort((a, b) => {
      if (a[1] !== b[1]) {
        return b[1] - a[1];
      }
      return a[0].localeCompare(b[0]);
    })
    .map(([word, count]) => `${word}: ${count}`);

  console.log(sortedWords.join('\n'));
});