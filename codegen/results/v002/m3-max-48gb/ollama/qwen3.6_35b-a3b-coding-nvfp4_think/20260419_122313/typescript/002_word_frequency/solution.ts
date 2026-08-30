import fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf-8');
const words = content.toLowerCase().match(/[a-z]+/g) || [];

const freqMap = new Map<string, number>();
for (const word of words) {
  freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

const entries = Array.from(freqMap.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

process.stdout.write(entries.map(([word, count]) => `${word}: ${count}`).join('\n') + '\n');