import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const normalized = content
  .toLowerCase()
  .replace(/[^a-z\s]/g, '');

const words = normalized.split(/\s+/).filter(word => word.length > 0);

const frequency: Map<string, number> = new Map();
for (const word of words) {
  frequency.set(word, (frequency.get(word) || 0) + 1);
}

const sorted = Array.from(frequency.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
  });

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}