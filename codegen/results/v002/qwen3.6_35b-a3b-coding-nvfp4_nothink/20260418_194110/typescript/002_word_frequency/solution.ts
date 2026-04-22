import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

// Convert to lowercase
const lowerContent = content.toLowerCase();

// Split into words: split by non-letter characters
// Replace all non-letter characters with spaces, then split by whitespace
const cleaned = lowerContent.replace(/[^a-z]/g, ' ');
const words = cleaned.split(/\s+/).filter(word => word.length > 0);

// Count frequencies
const freq: Record<string, number> = {};
for (const word of words) {
  freq[word] = (freq[word] || 0) + 1;
}

// Get unique words
const uniqueWords = Object.keys(freq);

// Sort: by count descending, then alphabetically ascending
uniqueWords.sort((a, b) => {
  const countDiff = freq[b] - freq[a];
  if (countDiff !== 0) return countDiff;
  return a.localeCompare(b);
});

// Output
for (const word of uniqueWords) {
  console.log(`${word}: ${freq[word]}`);
}