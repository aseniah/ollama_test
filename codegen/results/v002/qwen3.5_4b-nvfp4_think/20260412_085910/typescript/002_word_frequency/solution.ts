import { readFileSync } from 'fs';

const content = readFileSync('input/text.txt', 'utf8');

// Split into words, strip punctuation (keep only letters), lowercase
const words = content
  .match(/([a-z0-9]+)/g) ?? []
  .map(word => word.toLowerCase());

if (!words.length) {
  process.exit(0);
}

const freqMap: Map<string, number> = new Map();

for (const word of words) {
  const count = freqMap.get(word) ?? 0;
  freqMap.set(word, count + 1);
}

// Convert to array and sort by count descending, then alphabetically ascending
const freqArray = Array.from(freqMap.entries());
freqArray.sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of freqArray) {
  console.log(`${word}: ${count}`);
}