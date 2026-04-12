import * as fs from 'fs';
import * as path from 'path';

const filePath = process.argv[2] || 'input/text.txt';

try {
  const text = fs.readFileSync(filePath, 'utf8');

  // Convert to lowercase
  const lowerText = text.toLowerCase();

  // Strip all punctuation, keep only letters
  const cleanText = lowerText.replace(/[^a-z]/g, ' ');

  // Split into words and filter empty strings
  const words = cleanText.split(/\s+/).filter(word => word.length > 0);

  // Count frequency
  const frequencyMap = new Map<string, number>();
  for (const word of words) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }

  // Convert to array and sort
  const sortedEntries = Array.from(frequencyMap.entries())
    .sort((a, b) => {
      // Sort by count descending
      if (b[1] !== a[1]) {
        return b[1] - a[1];
      }
      // For ties, sort alphabetically ascending
      return a[0].localeCompare(b[0]);
    });

  // Output
  for (const [word, count] of sortedEntries) {
    console.log(`${word}: ${count}`);
  }
} catch (error) {
  console.error('Error reading file:', error);
  process.exit(1);
}