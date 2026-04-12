import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'text.txt');

try {
  const content = fs.readFileSync(filePath, 'utf-8');
  
  // Normalize to lowercase
  const lowerContent = content.toLowerCase();
  
  // Split into tokens based on non-letter characters
  // This regex matches sequences of letters, stripping all punctuation
  const tokens = lowerContent.match(/[a-z]+/g) || [];
  
  // Count frequencies
  const frequencyMap = new Map<string, number>();
  
  for (const word of tokens) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }
  
  // Convert to array for sorting
  const entries = Array.from(frequencyMap.entries());
  
  // Sort: by count descending, then alphabetically ascending for ties
  entries.sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Descending count
    }
    return a[0].localeCompare(b[0]); // Ascending alphabetically
  });
  
  // Output results
  for (const [word, count] of entries) {
    console.log(`${word}: ${count}`);
  }
  
} catch (error) {
  console.error('Error reading file:', error);
  process.exit(1);
}