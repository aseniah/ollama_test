import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a file from 'input/text.txt', processes the text to count
 * word frequencies based on specific rules, and outputs the result to stdout.
 */

function solve() {
    const filePath = path.join(process.cwd(), 'input', 'text.txt');

    let content: string;

    try {
        content = fs.readFileSync(filePath, 'utf-8');
    } catch (err) {
        // If file doesn't exist or can't be read, exit silently as per "Do not output anything else"
        process.exit(0);
    }

    // 1. Convert to lowercase
    const lowercased = content.toLowerCase();

    // 2. Strip punctuation (keep only letters and whitespace for splitting)
    // We replace everything that isn't a lowercase letter or whitespace with a space
    // to ensure words separated by punctuation (like "word.word") are split correctly.
    const cleaned = lowercased.replace(/[^a-z\s]/g, ' ');

    // 3. Split into words and filter out empty strings
    const words = cleaned.split(/\s+/).filter(word => word.length > 0);

    // 4. Count frequency
    const frequencyMap: Map<string, number> = new Map();

    for (const word of words) {
        const count = frequencyMap.get(word) || 0;
        frequencyMap.set(word, count + 1);
    }

    // 5. Convert map to array for sorting
    const sortedEntries = Array.from(frequencyMap.entries());

    // 6. Sort: Count descending, then Alphabetical ascending
    sortedEntries.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA;
        }
        
        // Tie-breaker: Alphabetical ascending
        if (wordA < wordB) return -1;
        if (wordA > wordB) return 1;
        return 0;
    });

    // 7. Output results
    for (const [word, count] of sortedEntries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();