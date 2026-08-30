import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const words = content.toLowerCase().split(/[^a-z]+/).filter(w => w.length > 0);

const freq: Map<string, number> = new Map();
for (const word of words) {
  freq.set(word, (freq.get(word) ?? 0) + 1);
}

const entries = [...freq.entries()];
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}