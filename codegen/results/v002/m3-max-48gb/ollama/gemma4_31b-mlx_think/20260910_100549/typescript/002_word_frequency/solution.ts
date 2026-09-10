import * as fs from 'fs';
import * as path from 'path';

/**
 * Solves the word frequency counting task.
 * Reads from 'input/text.txt', processes words based on specified rules,
 * and outputs the frequency sorted by count descending and alphabetically ascending.
 */
function solve() {
    const inputPath = path.join('input', 'text.txt');
    
    try {
        // Read the file content
        const content = fs.readFileSync(inputPath, 'utf8');

        // 1. Convert all words to lowercase
        // 2. Strip all punctuation (keep only letters)
        // We split the text by any character that is not a lowercase letter (a-z).
        // This effectively removes punctuation and treats them as delimiters.
        const words = content
            .toLowerCase()
            .split(/[^a-z]+/)
            .filter(word => word.length > 0);

        // Count the frequency of each word
        const frequencyMap = new Map<string, number>();
        for (const word of words) {
            frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
        }

        // Convert Map entries to an array for sorting
        const sortedEntries = Array.from(frequencyMap.entries());

        // Sort by count descending; for ties, sort alphabetically ascending
        sortedEntries.sort((a, b) => {
            const countA = a[1];
            const countB = b[1];
            const wordA = a[0];
            const wordB = b[0];

            if (countB !== countA) {
                return countB - countA; // Primary sort: count descending
            }
            return wordA.localeCompare(wordB); // Secondary sort: word ascending
        });

        // Output in the format: word: count
        for (const [word, count] of sortedEntries) {
            process.stdout.write(`${word}: ${count}\n`);
        }
    } catch (err) {
        // The prompt asks for specific output; we handle potential file errors silently 
        // or based on standard Node.js failure unless specified otherwise.
    }
}

// Execute the solver
solve();