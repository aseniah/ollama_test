import * as fs from 'fs';
import * as path from 'path';

// Helper function to sanitize words by converting to lowercase and removing punctuation
function sanitizeWord(word: string): string {
    return word.toLowerCase().replace(/[^a-z]/g, '');
}

// Read the file asynchronously
fs.readFile(path.resolve(__dirname, 'input', 'text.txt'), 'utf8', (err, data) => {
    if (err) {
        console.error('Error reading file:', err);
        return;
    }

    // Split the text into words and sanitize them
    const words = data.split(/\s+/).map(sanitizeWord);

    // Create a map to count word frequencies
    const wordCount: { [key: string]: number } = {};

    for (const word of words) {
        if (word) {
            wordCount[word] = (wordCount[word] || 0) + 1;
        }
    }

    // Sort the words by frequency and alphabetically for ties
    const sortedWords = Object.entries(wordCount).sort((a, b) => {
        if (b[1] === a[1]) {
            return a[0].localeCompare(b[0]);
        }
        return b[1] - a[1];
    });

    // Output the results
    for (const [word, count] of sortedWords) {
        console.log(`${word}: ${count}`);
    }
});