import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const filePath = path.resolve('input/text.txt');
let content: string;
try {
    content = fs.readFileSync(filePath, 'utf-8');
} catch (e) {
    // If the file doesn't exist, exit
    process.exit(1);
}

// Split into words, convert to lowercase, strip punctuation (keep only letters)
const words = content
    .toLowerCase()
    .split(/\s+/)
    .map(word => word.replace(/[^a-z]/g, ''))
    .filter(word => word.length > 0);

// Count frequency
const freq: Map<string, number> = new Map();
for (const word of words) {
    freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array and sort: by count descending, then alphabetically ascending
const entries = Array.from(freq.entries());
entries.sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    // Alphabetically ascending
    return a[0].localeCompare(b[0]);
});

// Output one line per unique word
for (const [word, count] of entries) {
    console.log(`${word}: ${count}`);
}