import { readFileSync } from 'fs';

const text = readFileSync('input/text.txt', 'utf8');

const freq = new Map<string, number>();

for (const raw of text.split(/\s+/)) {
  const word = raw.toLowerCase().replace(/[^a-z]/g, '');
  if (word.length > 0) {
    freq.set(word, (freq.get(word) ?? 0) + 1);
  }
}

const sorted = [...freq.entries()].sort(([a, ca], [b, cb]) =>
  cb !== ca ? cb - ca : a.localeCompare(b)
);

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}
