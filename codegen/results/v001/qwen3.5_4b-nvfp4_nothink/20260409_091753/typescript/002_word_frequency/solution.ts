import { readFileSync } from "fs";
import { dirname, join } from "path";

const inputPath = process.argv[2];
if (!inputPath) {
  console.error("Usage: node script.ts <filepath>");
  process.exit(1);
}

// Normalize path to be absolute so it works correctly with Node.js
const absolutePath = join(dirname(inputPath), "input", "text.txt");

const content = readFileSync(absolutePath, "utf-8");

// Convert to lowercase and replace anything that's not a letter with a space
const normalized = content.toLowerCase().replace(/[^a-z]/g, " ");

const words: string[] = normalized.split(" ").filter((w) => w.length > 0);

const freqMap = new Map<string, number>();

for (const word of words) {
  const currentCount = freqMap.get(word) || 0;
  freqMap.set(word, currentCount + 1);
}

// Sort: descending by count, then ascending alphabetically for ties
const sortedWords = Array.from(freqMap.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
  })
  .map(([word, count]) => `${word}: ${count}`);

// Output one line per word
console.log(sortedWords.join("\n"));