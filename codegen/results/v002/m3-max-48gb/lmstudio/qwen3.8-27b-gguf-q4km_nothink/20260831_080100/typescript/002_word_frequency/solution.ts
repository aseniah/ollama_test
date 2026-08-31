import * as fs from 'fs';

// Read the input file
const text = fs.readFileSync('input/text.txt', 'utf-8');

// Process the text
// Convert to lowercase
const lower = text.toLowerCase();

// Strip all non-letter characters, replacing with spaces to preserve word boundaries
// Then split on whitespace
const cleaned = lower.replace(/[^a-z]/g, ' ');

// Split into words
const words = cleaned.split(/\s+/).filter(w => w.length > 0);

// Count frequencies
const freq: Map<string, number> = new Map();
for (const word of words) {
    freq.set(word, (freq.get(word) || 0) + 1);
}

// Convert to array of [word, count]
const entries: [string, number][] = Array.from(freq.entries());

// Sort by count descending, then alphabetically ascending
entries.sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
if (output.length > 0) {
    process.stdout.write(output + '\n');
}