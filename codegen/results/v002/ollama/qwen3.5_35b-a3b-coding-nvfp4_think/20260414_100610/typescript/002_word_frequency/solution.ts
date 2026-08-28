import * as fs from 'fs';
import * as path from 'path';

// Read the file
const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

// Split into words and process
const words = content.split(/\s+/);

// Count frequencies
const wordCount = new Map<string, number>();

for (const word of words) {
    // Convert to lowercase and strip non-letter characters
    const cleaned = word.toLowerCase().replace(/[^a-z]/g, '');
    
    // Only count non-empty words
    if (cleaned.length > 0) {
        wordCount.set(cleaned, (wordCount.get(cleaned) || 0) + 1);
    }
}

// Convert to array and sort
const sorted = Array.from(wordCount.entries())
    .sort((a, b) => {
        // Sort by count descending
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        // Then alphabetically ascending
        return a[0].localeCompare(b[0]);
    });

// Output
for (const [word, count] of sorted) {
    console.log(`${word}: ${count}`);
}