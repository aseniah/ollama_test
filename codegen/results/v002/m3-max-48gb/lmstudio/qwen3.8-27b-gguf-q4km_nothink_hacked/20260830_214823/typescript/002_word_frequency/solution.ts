import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.join(__dirname, 'input', 'text.txt');
let content: string;

try {
  content = fs.readFileSync(inputPath, 'utf-8');
} catch (e) {
  // If the file doesn't exist in the expected location, try current directory
  content = fs.readFileSync('input/text.txt', 'utf-8');
}

// Process the content: convert to lowercase, strip punctuation (keep only letters)
const lowercased = content.toLowerCase();

// Keep only letters and spaces (or other non-punctuation characters), then extract words
// Strip all punctuation - keep only letters
const cleanText = lowercased.replace(/[^a-z\s]/g, ' ');

// Split into words
const words = cleanText.split(/\s+/).filter(w => w.length > 0);

// Count frequency
const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array for sorting
const entries: [string, number][] = Array.from(freq.entries());

// Sort by count descending, then alphabetically ascending for ties
entries.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}