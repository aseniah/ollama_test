import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'text.txt');

try {
  const content = fs.readFileSync(filePath, 'utf-8');
  
  // Convert to lowercase and extract words (only letters)
  const words = content
    .toLowerCase()
    .match(/[a-z]+/g) || [];
  
  // Count word frequencies
  const wordCounts: { [key: string]: number } = {};
  for (const word of words) {
    wordCounts[word] = (wordCounts[word] || 0) + 1;
  }
  
  // Sort by count (descending) then by word (ascending)
  const sortedWords = Object.keys(wordCounts).sort((a, b) => {
    if (wordCounts[b] !== wordCounts[a]) {
      return wordCounts[b] - wordCounts[a]; // Descending count
    }
    return a.localeCompare(b); // Ascending alphabetical
  });
  
  // Output results
  for (const word of sortedWords) {
    console.log(`${word}: ${wordCounts[word]}`);
  }
} catch (error) {
  // If file doesn't exist or other error, output nothing
}