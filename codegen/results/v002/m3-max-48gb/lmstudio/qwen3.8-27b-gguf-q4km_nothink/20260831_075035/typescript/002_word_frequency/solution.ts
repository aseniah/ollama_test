import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.join(__dirname, 'input', 'text.txt');
let text: string;

try {
    text = fs.readFileSync(inputPath, 'utf-8');
} catch (e) {
    // Fallback: try relative to current working directory
    text = fs.readFileSync(path.join(process.cwd(), 'input', 'text.txt'), 'utf-8');
}

// Process the text
const words = text
    .toLowerCase()
    .replace(/[^a-z\s]/g, ' ')  // Strip all punctuation, keep only letters and whitespace
    .split(/\s+/)               // Split by whitespace
    .filter(word => word.length > 0); // Remove empty strings

// Count frequency
const frequency: Map<string, number> = new Map();
for (const word of words) {
    frequency.set(word, (frequency.get(word) || 0) + 1);
}

// Convert to array for sorting
const entries = Array.from(frequency.entries());

// Sort by count descending, then alphabetically ascending
entries.sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
console.log(output);