import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8').toLowerCase();
const words = content.match(/[a-z]+/g) || [];

const freqMap = new Map<string, number>();
for (const word of words) {
  freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

const entries = Array.from(freqMap.entries());
entries.sort((a, b) => {
  // Sort by count descending
  if (b[1] !== a[1]) return b[1] - a[1];
  // For ties, sort alphabetically ascending
  return a[0] < b[0] ? -1 : 1;
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}