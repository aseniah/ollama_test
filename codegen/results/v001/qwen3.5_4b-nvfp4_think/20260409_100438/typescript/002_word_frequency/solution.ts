import { readFileSync } from "fs";
import { resolve, join } from "path";

const filePath = resolve(join(process.cwd(), "input", "text.txt"));
const content = readFileSync(filePath, "utf-8");

const wordsMap = new Map<string, number>();
const punctuationPattern = /[^\w]/gu;

for (const char of content) {
  if (!char.match(/[a-z]/i)) continue;
}

const text = content.toLowerCase().match(/[\w\s]+/g)?.join("") || "";
const lowercaseText = text.replace(punctuationPattern, "");

const words = lowercaseText.split(/\s+/).filter(word => word.length > 0);

for (const word of words) {
  const cleanedWord = word.match(/[a-z]/gu)?.join("") || "";
  if (cleanedWord && cleanedWord.length > 0) {
    const currentCount = wordsMap.get(cleanedWord) || 0;
    wordsMap.set(cleanedWord, currentCount + 1);
  }
}

const entries = Array.from(wordsMap.entries());
entries.sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}