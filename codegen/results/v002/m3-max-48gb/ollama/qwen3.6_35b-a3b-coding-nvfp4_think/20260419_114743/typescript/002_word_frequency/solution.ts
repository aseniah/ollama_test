import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const words = content.toLowerCase().match(/[a-z]+/g) || [];
const freq: Record<string, number> = {};

for (const word of words) {
  freq[word] = (freq[word] || 0) + 1;
}

const sortedWords = Object.entries(freq).sort((a, b) => {
  if (b[1] !== a[1]) {
    return b[1] - a[1];
  }
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}