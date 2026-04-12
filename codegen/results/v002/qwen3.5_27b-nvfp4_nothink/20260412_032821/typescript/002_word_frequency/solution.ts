import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.resolve(process.cwd(), 'input', 'text.txt');

try {
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  // Convert to lowercase and extract words (letters only)
  const words = content.toLowerCase().match(/[a-z]+/g) || [];
  
  const frequencyMap: Map<string, number> = new Map();
  
  // Count frequencies
  for (const word of words) {
    if (word.length > 0) {
      frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }
  }
  
  // Convert to array and sort
  const sortedEntries = Array.from(frequencyMap.entries())
    .sort((a, b) => {
      if (b[1] !== a[1]) {
        return b[1] - a[1]; // Sort by count descending
      }
      return a[0].localeCompare(b[0]); // Sort alphabetically ascending for ties
    });
  
  // Output results
  for (const [word, count] of sortedEntries) {
    console.log(`${word}: ${count}`);
  }
} catch (error) {
  process.exit(1);
}