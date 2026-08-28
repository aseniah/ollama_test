const fs = require('fs');
const path = require('path');

const filePath = path.resolve('input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const words = content
  .toLowerCase()
  .split(/\s+/)
  .map((w: string) => w.replace(/[^a-z]/g, ''))
  .filter((w: string) => w.length > 0);

const freq: Record<string, number> = {};
for (const word of words) {
  freq[word] = (freq[word] || 0) + 1;
}

const sorted = Object.entries(freq).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}
