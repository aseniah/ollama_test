import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'text.txt');

try {
  const data = fs.readFileSync(filePath, 'utf8');
  const words = data.toLowerCase().match(/[a-z]+/g) || [];
  
  const wordCount: { [key: string]: number } = {};
  
  words.forEach(word => {
    wordCount[word] = (wordCount[word] || 0) + 1;
  });
  
  const sortedWords = Object.entries(wordCount)
    .sort(([wordA, countA], [wordB, countB]) => {
      if (countB !== countA) {
        return countB - countA; // Sort by count descending
      }
      return wordA.localeCompare(wordB); // Sort alphabetically ascending for ties
    });
  
  sortedWords.forEach(([word, count]) => {
    console.log(`${word}: ${count}`);
  });
} catch (error) {
  // Do nothing if file doesn't exist or other error
}