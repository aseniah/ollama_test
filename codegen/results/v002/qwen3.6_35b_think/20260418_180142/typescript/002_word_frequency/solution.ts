import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const words = content.split(/\s+/);
const freq = new Map<string, number>();

for (const word of words) {
  const lower = word.toLowerCase();
  const cleaned = lower.replace(/[^a-z]/g, '');
  if (cleaned.length > 0) {
    freq.set(cleaned, (freq.get(cleaned) || 0) + 1);
  }
}

const result = Array.from(freq.entries()).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0] < b[0] ? -1 : 1;
});

for (const [word, count] of result) {
  console.log(`${word}: ${count}`);
}