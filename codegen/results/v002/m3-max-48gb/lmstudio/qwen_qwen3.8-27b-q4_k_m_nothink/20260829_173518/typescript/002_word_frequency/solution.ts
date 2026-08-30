import * as fs from 'fs';

// Read the file
const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Process the content
// Convert to lowercase
const lowerContent = content.toLowerCase();

// Extract words: keep only letters
const words = lowerContent.match(/[a-z]+/g);

if (!words) {
    process.exit(0);
}

// Count frequency
const freqMap = new Map<string, number>();
for (const word of words) {
    freqMap.set(word, (freqMap.get(word) || 0) + 1);
}

// Convert to array of [word, count]
const entries = Array.from(freqMap.entries());

// Sort by count descending; for ties, sort alphabetically ascending
entries.sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    // Alphabetically ascending
    return a[0].localeCompare(b[0]);
});

// Output
const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
if (output) {
    process.stdout.write(output + '\n');
}