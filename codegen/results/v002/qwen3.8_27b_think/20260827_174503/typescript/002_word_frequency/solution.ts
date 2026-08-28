import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

const freq = new Map<string, number>();
for (const word of words) {
  freq.set(word, (freq.get(word) ?? 0) + 1);
}

const entries = Array.from(freq.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}