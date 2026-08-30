import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a file from 'input/text.txt', processes the text to count
 * word frequencies according to specific rules, and outputs the results.
 */

function solve() {
    const inputPath = path.join('input', 'text.txt');

    let content: string;

    try {
        // Read file synchronously
        content = fs.readFileSync(inputPath, 'utf-8');
    } catch (err) {
        // If file does not exist or cannot be read, exit silently as per requirements
        return;
    }

    // 1. Convert all words to lowercase
    const lowercased = content.toLowerCase();

    // 2. Strip all punctuation (keep only letters)
    // We replace anything that is not a lowercase letter or whitespace with an empty string.
    // We also replace punctuation with spaces to ensure words aren't accidentally merged.
    const cleaned = lowercased.replace(/[^a-z\s]/g, ' ');

    // Split by whitespace and filter out empty strings resulting from multiple spaces
    const words = cleaned.split(/\s+/).filter(word => word.length > 0);

    // 3. Count frequencies
    const frequencyMap: Record<string, number> = {};
    for (const word of words) {
        frequencyMap[word] = (frequencyMap[word] || 0) + 1;
    }

    // Convert map to array of [word, count] pairs for sorting
    const sortedEntries = Object.entries(frequencyMap).sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        // Sort by count descending
        if (countB !== countA) {
            return countB - countA;
        }

        // For ties, sort alphabetically ascending
        return wordA.localeCompare(wordB);
    });

    // 4. Output results in format 'word: count'
    for (const [word, count] of sortedEntries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

// Run the solver
solve();