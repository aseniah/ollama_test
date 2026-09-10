import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a text file, counts the frequency of each word,
 * and outputs the result sorted by frequency descending and alphabetically ascending.
 */
function solve() {
    const inputPath = path.join('input', 'text.txt');

    try {
        // Read the file content
        const content = fs.readFileSync(inputPath, 'utf8');

        // Process the text:
        // 1. Convert to lowercase
        // 2. Replace all non-letter characters with spaces to treat them as word boundaries
        // 3. Split by whitespace and filter out empty strings
        const words = content
            .toLowerCase()
            .replace(/[^a-z\s]/g, ' ')
            .split(/\s+/)
            .filter(word => word.length > 0);

        // Map to store word frequencies
        const wordCounts = new Map<string, number>();
        for (const word of words) {
            wordCounts.set(word, (wordCounts.get(word) || 0) + 1);
        }

        // Convert Map to an array of [word, count] pairs for sorting
        const sortedWords = Array.from(wordCounts.entries()).sort((a, b) => {
            const [wordA, countA] = a;
            const [wordB, countB] = b;

            // Sort by count descending
            if (countB !== countA) {
                return countB - countA;
            }
            // For ties, sort alphabetically ascending
            return wordA.localeCompare(wordB);
        });

        // Print each unique word and its count in the format "word: count"
        for (const [word, count] of sortedWords) {
            process.stdout.write(`${word}: ${count}\n`);
        }
    } catch (err) {
        // Handle file not found or read errors silently as per prompt "Do not output anything else"
        // however, in a real app, we would log this.
    }
}

// Execute the solve function
solve();