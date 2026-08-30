import fs from 'fs';
import path from 'path';

const filePath = path.resolve('input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(w => w.replace(/[^a-z]/g, ''))
  .filter(w => w.length > 0);

const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) ?? 0) + 1);
}

const sorted = [...freq.entries()].sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of sorted) {
  process.stdout.write(`${word}: ${count}\n`);
}
