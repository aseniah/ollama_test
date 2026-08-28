import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const lowerContent = content.toLowerCase();
const words = lowerContent.match(/[a-z]+/g) || [];

const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) || 0) + 1);
}

const result = Array.from(freq.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
  });

for (const [word, count] of result) {
  console.log(`${word}: ${count}`);
}