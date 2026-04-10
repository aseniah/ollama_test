import * as fs from 'fs';

// Check if a filename is provided via arguments, otherwise default to the task specific file
const filePath = process.argv[2] || 'input/text.txt';

try {
    // Read the file content synchronously
    const text = fs.readFileSync(filePath, 'utf-8');

    // Normalize: convert to lowercase
    // Extract words: keep only sequences of letters [a-z]
    const words = text.toLowerCase().match(/[a-z]+/g) || [];

    // Count frequency of each word
    const counts = new Map<string, number>();
    for (const word of words) {
        counts.set(word, (counts.get(word) || 0) + 1);
    }

    // Convert map to array and sort
    // Sort rules: 
    // 1. Count descending (b[1] - a[1])
    // 2. Word ascending alphabetically (a[0].localeCompare(b[0]))
    const sortedEntries = Array.from(counts).sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });

    // Format output: "word: count" joined by newline
    const output = sortedEntries.map(([word, count]) => `${word}: ${count}`).join('\n');

    // Output to stdout
    if (output) {
        console.log(output);
    }
} catch (error) {
    // Exit silently on failure
    process.exit(1);
}