import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.join(process.cwd(), 'input', 'text.txt');

try {
  const content = fs.readFileSync(inputFile, 'utf-8');
  
  // Split by whitespace and filter out empty strings
  const rawWords = content.trim().split(/\s+/).filter(word => word.length > 0);
  
  // Process words: convert to lowercase and strip non-letters
  const processedWords = rawWords.map(word => {
    return word.toLowerCase().replace(/[^a-z]/g, '');
  }).filter(word => word.length > 0);

  // Count frequencies
  const frequencyMap = new Map<string, number>();
  
  for (const word of processedWords) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }

  // Convert to array for sorting
  const wordCounts = Array.from(frequencyMap.entries());

  // Sort: descending by count, then ascending alphabetically
  wordCounts.sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1]; // Descending count
    }
    return a[0].localeCompare(b[0]); // Ascending alphabetical
  });

  // Output results
  for (const [word, count] of wordCounts) {
    console.log(`${word}: ${count}`);
  }
} catch (error) {
  // If file not found or error, output nothing as per "Do not output anything else"
  // But in a real scenario, we might want to handle this differently. 
  // Given the strict rules, we do nothing here.
}