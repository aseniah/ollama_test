import * as fs from 'fs';
import * as path from 'path';

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'text.txt');
  
  if (!fs.existsSync(inputPath)) {
    return;
  }

  const content = fs.readFileSync(inputPath, 'utf-8');
  
  // Split by non-letter characters, filter empty strings, convert to lowercase
  const words = content
    .toLowerCase()
    .replace(/[^a-z\s]/g, '')
    .split(/\s+/)
    .filter(word => word.length > 0);

  const frequencyMap = new Map<string, number>();

  for (const word of words) {
    frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
  }

  const entries = Array.from(frequencyMap.entries());

  // Sort: count descending, then word alphabetically ascending
  entries.sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
  });

  for (const [word, count] of entries) {
    console.log(`${word}: ${count}`);
  }
}

main();