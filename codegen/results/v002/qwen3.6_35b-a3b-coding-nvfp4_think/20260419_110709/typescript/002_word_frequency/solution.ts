import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const words = content.toLowerCase().match(/[a-z]+/g) || [];

const counts: Record<string, number> = {};
for (const word of words) {
  counts[word] = (counts[word] || 0) + 1;
}

const sorted = Object.entries(counts).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}