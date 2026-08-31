import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.resolve(__dirname, '../input/text.txt');

let content: string;
try {
    content = fs.readFileSync(inputPath, 'utf-8');
} catch {
    // Try relative to cwd
    try {
        content = fs.readFileSync('input/text.txt', 'utf-8');
    } catch {
        process.stdout.write('Error: Could not read input file\n');
        process.exit(1);
    }
}

// Strip punctuation: keep only letters (a-z), then lowercase
const words = content
    .toLowerCase()
    .replace(/[^a-z\s]/g, '')
    .split(/\s+/)
    .filter(w => w.length > 0);

// Count frequencies
const freqMap = new Map<string, number>();
for (const word of words) {
    freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

// Sort by count descending, then alphabetically ascending
const sorted = [...freqMap.entries()].sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
});

// Output
const lines = sorted.map(([word, count]) => `${word}: ${count}`);
process.stdout.write(lines.join('\n') + '\n');