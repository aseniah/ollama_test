import * as fs from "fs";
import * as path from "path";

const inputPath = process.argv[2] || "input/text.txt";
const content = fs.readFileSync(inputPath, "utf8");

const words = content
  .toLowerCase()
  .replace(/[^a-z0-9]+/g, " ")
  .split(" ")
  .filter((w) => w.length > 0);

const freqMap = new Map<string, number>();
for (const word of words) {
  freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

const sorted = [...freqMap.entries()].sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of sorted) {
  console.log(`${word}: ${count}`);
}