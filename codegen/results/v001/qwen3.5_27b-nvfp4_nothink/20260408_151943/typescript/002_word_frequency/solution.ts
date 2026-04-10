import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'text.txt');

try {
  const content = fs.readFileSync(inputPath, 'utf8');
  // Split by non-alphabetic characters to strip punctuation and get words
  const rawWords = content.split(/[^a-zA-Z]+/);
  
  const wordCounts = new Map<string, number>();

  for (const rawWord of rawWords) {
    const word = rawWord.toLowerCase();
    if (word.length === 0) continue; // Skip empty strings resulting from split
    
    wordCounts.set(word, (wordCounts.get(word) || 0) + 1);
  }

  // Convert map to array for sorting
  const sortedWords: { word: string; count: number }[] = Array.from(wordCounts.entries())
    .map(([word, count]) => ({ word, count }))
    .sort((a, b) => {
      if (b.count !== a.count) {
        return b.count - a.count; // Sort by count descending
      }
      return a.word.localeCompare(b.word); // Sort alphabetically ascending for ties
    });

  // Output results
  sortedWords.forEach(({ word, count }) => {
    console.log(`${word}: ${count}`);
  });
} catch (error) {
  // Do not output anything else on error as per instructions
}