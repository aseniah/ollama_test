import fs from "fs";

const filePath = process.argv[2];
if (!filePath) {
  console.error("Usage: tsx solution.ts <input-file>");
  process.exit(1);
}

const content = fs.readFileSync(filePath, "utf-8");
const lines = content.split("\n");

// Map to store word frequencies
const wordCountMap = new Map<string, number>();

for (const line of lines) {
  let text = line.toLowerCase();

  // Remove non-alphabetic characters
  text = text.replace(/[^a-z]/g, " ");

  const words = text.split(/\s+/);

  for (const word of words) {
    if (word.length > 0) {
      const count = wordCountMap.get(word) || 0;
      wordCountMap.set(word, count + 1);
    }
  }
}

// Convert to array and sort: primary by count (desc), secondary by word (asc)
const sortedWords = Array.from(wordCountMap.entries())
  .sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
  })
  .map(([word, count]) => [word, count]);

for (const [word, count] of sortedWords) {
  console.log(`${word}: ${count}`);
}