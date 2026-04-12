import * as fs from 'fs';
import * as path from 'path';

// Read the file content
const filePath = path.join(__dirname, 'input', 'text.txt');
let content: string;

try {
  content = fs.readFileSync(filePath, 'utf-8');
} catch (error) {
  // If the file cannot be read, exit silently or handle as needed per requirements.
  // The prompt implies the file exists, so we assume success for the logic part.
  process.exit(1);
}

// Helper to strip punctuation and convert to lowercase
function cleanWord(word: string): string {
  return word.toLowerCase().replace(/[^a-z]/g, '');
}

// Split content by non-letter characters to extract words
const rawWords = content.split(/[^a-zA-Z0-9]+/);

const frequencyMap: Map<string, number> = new Map();

for (const word of rawWords) {
  const clean = cleanWord(word);
  if (clean.length > 0) {
    const count = frequencyMap.get(clean) || 0;
    frequencyMap.set(clean, count + 1);
  }
}

// Convert map to array for sorting
const wordCounts: Array<{ word: string; count: number }> = [];
for (const [word, count] of frequencyMap.entries()) {
  wordCounts.push({ word, count });
}

// Sort: 
// 1. By count descending
// 2. Alphabetically ascending for ties
wordCounts.sort((a, b) => {
  if (b.count !== a.count) {
    return b.count - a.count; // Descending count
  }
  return a.word.localeCompare(b.word); // Ascending alphabet
});

// Output
for (const item of wordCounts) {
  console.log(`${item.word}: ${item.count}`);
}