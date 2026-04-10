import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input/text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const words = content
  .toLowerCase()
  .split(/\s+/)
  .map(word => word.replace(/[^a-z]/g, ''))
  .filter(word => word.length > 0);

const frequency: Record<string, number> = {};

for (const word of words) {
  frequency[word] = (frequency[word] || 0) + 1;
}

const sorted = Object.entries(frequency)
  .sort((a, b) => {
    if (b[1] !== a[1]) {
      return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
  });

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}
