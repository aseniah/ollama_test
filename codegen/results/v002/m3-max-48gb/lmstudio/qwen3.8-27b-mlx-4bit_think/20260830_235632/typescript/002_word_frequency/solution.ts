import * as fs from 'fs';

const text = fs.readFileSync('input/text.txt', 'utf-8');
const words = text
  .toLowerCase()
  .split(/\s+/)
  .map(w => w.replace(/[^a-z]/g, ''))
  .filter(w => w.length > 0);

const freq: Record<string, number> = {};
for (const word of words) {
  freq[word] = (freq[word] || 0) + 1;
}

const entries = Object.entries(freq).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}