import { readFileSync } from "node:fs";
import * as path from "node:path";

const inputPath = path.join(process.cwd(), "input", "text.txt");
const content = readFileSync(inputPath, "utf-8");

const words: string[] = [];

for (let i = 0; i < content.length; ) {
  const char = content[i];
  
  if (/[a-z]/.test(char)) {
    let wordStart = i;
    while (i < content.length && /[a-z]/.test(content[i])) {
      i++;
    }
    words.push(content.substring(wordStart, i).toLowerCase());
  }
  i++;
}

const freqMap: Map<string, number> = new Map();

for (const word of words) {
  if (!freqMap.has(word)) {
    freqMap.set(word, 0);
  }
  freqMap.set(word, freqMap.get(word)! + 1);
}

const sortedWords = Array.from(freqMap.entries())
  .sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))
  .map(([word, count]) => `${word}: ${count}`);

console.log(sortedWords.join("\n"));