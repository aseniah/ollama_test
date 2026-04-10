import * as fs from 'fs';

// Default path is input/text.txt, can be overridden by CLI argument
const filePath = process.argv[2] || 'input/text.txt';

try {
    // Read file content
    const text = fs.readFileSync(filePath, 'utf-8');

    // Convert to lowercase and extract words (only letters a-z)
    // match returns null if no matches found, so we default to empty array
    const words = text.toLowerCase().match(/[a-z]+/g) || [];

    // Count frequencies using a Map
    const counts: Map<string, number> = new Map();
    for (const word of words) {
        counts.set(word, (counts.get(word) || 0) + 1);
    }

    // Convert map entries to array and sort
    const sortedEntries = Array.from(counts).sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1]; // Sort by count descending
        }
        return a[0].localeCompare(b[0]); // Sort by word ascending alphabetically
    });

    // Output results
    for (const [word, count] of sortedEntries) {
        console.log(`${word}: ${count}`);
    }
} catch (error) {
    process.exit(1);
}