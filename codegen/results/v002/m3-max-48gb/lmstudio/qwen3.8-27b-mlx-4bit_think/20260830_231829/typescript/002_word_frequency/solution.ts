import * as fs from "fs";

const text = fs.readFileSync("input/text.txt", "utf-8");

// Split into tokens by whitespace, strip non-letters, lowercase, skip empty
const words = text
  .split(/\s+/)
  .map((w) => w.replace(/[^a-zA-Z]/g, "").toLowerCase())
  .filter((w) => w.length > 0);

// Count frequencies
const freq: Record<string, number> = {};
for (const word of words) {
  freq[word] = (freq[word] || 0) + 1;
}

// Sort: count descending, then alphabetically ascending
const entries = Object.entries(freq).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1];
  return a[0].localeCompare(b[0]);
});

// Output
for (const [word, count] of entries) {
  console.log(`${word}: ${count}`);
}