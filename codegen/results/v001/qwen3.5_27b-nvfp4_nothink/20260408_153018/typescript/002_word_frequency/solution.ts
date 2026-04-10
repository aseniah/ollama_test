import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.join('input', 'text.txt');

try {
  const content = fs.readFileSync(inputFile, 'utf8');
  
  // Split by whitespace/newlines, convert to lowercase, strip non-letters
  const words = content
    .toLowerCase()
    .split(/\s+/)
    .map(word => word.replace(/[^a-z]+/g, ''))
    .filter(word => word.length > 0);

  const frequencyMap: Map<string, number> = new Map();

  for (const word of words) {
    const count = frequencyMap.get(word) || 0;
    frequencyMap.set(word, count + 1);
  }

  // Convert to array and sort
  const sortedWords = Array.from(frequencyMap.entries())
    .sort((a, b) => {
      if (b[1] !== a[1]) {
        return b[1] - a[1]; // Descending by count
      }
      return a[0].localeCompare(b[0]); // Ascending alphabetically for ties
    });

  // Output results
  for (const [word, count] of sortedWords) {
    console.log(`${word}: ${count}`);
  }
} catch (error) {
  // If file not found or other error, do nothing as per "Do not output anything else"
}