import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(__dirname, 'input', 'text.txt');

let content: string;
try {
    content = fs.readFileSync(inputPath, 'utf-8');
} catch (e) {
    // Fallback: try relative path
    content = fs.readFileSync('input/text.txt', 'utf-8');
}

// Split into words, convert to lowercase, strip non-letters
const words = content
    .toLowerCase()
    .split(/\s+/)
    .map(w => w.replace(/[^a-z]/g, ''))
    .filter(w => w.length > 0);

// Count frequencies
const freq: Map<string, number> = new Map();
for (const word of words) {
    freq.set(word, (freq.get(word) || 0) + 1);
}

// Sort: count descending, then alphabetically ascending
const sorted = Array.from(freq.entries()).sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
});

// Output
const lines = sorted.map(([word, count]) => `${word}: ${count}`);
console.log(lines.join('\n'));