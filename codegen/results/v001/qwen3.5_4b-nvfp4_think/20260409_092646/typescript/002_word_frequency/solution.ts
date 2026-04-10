import { readFileSync } from 'fs';

const filePath = process.argv.length > 2 ? process.argv[2] : 'input/text.txt';

try {
  const text = readFileSync(filePath, { encoding: 'utf-8' });
  
  // Convert to lowercase and extract words consisting of only letters
  const words = text.toLowerCase().match(/[a-z]+/g);
  
  const freqMap = new Map<string, number>();
  
  if (words) {
    for (const word of words) {
      const currentCount = freqMap.get(word) || 0;
      freqMap.set(word, currentCount + 1);
    }
  }
  
  // Convert entries to array
  const result: [string, number][] = Array.from(freqMap.entries());
  
  // Sort by count descending, then word ascending
  result.sort((a, b) => {
    if (a[1] !== b[1]) {
      return b[1] - a[1];
    } else {
      return a[0].localeCompare(b[0]);
    }
  });
  
  for (const [word, count] of result) {
    console.log(`${word}: ${count}`);
  }
} catch (err) {
  // Do not output anything else
  process.exit(1);
}