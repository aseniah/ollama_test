import * as fs from 'fs';
import * as path from 'path';

/**
 * Problem: Count word frequency in a file.
 * Rules:
 * - Lowercase all words
 * - Remove punctuation (keep only letters/numbers/whitespace)
 * - Output: word: count
 * - Sort: Count descending, then Alphabetical ascending
 */

function solve() {
    const inputPath = path.join(process.cwd(), 'input', 'text.txt');

    let content: string;

    try {
        content = fs.readFileSync(inputPath, 'utf-8');
    } catch (error) {
        // If file doesn't exist, we exit silently as per instructions to not output anything else
        return;
    }

    // 1. Convert to lowercase
    const lowerContent = content.toLowerCase();

    // 2. Strip punctuation: Replace anything that isn't a letter, number, or whitespace with space
    // We use a regex that keeps alphanumeric characters and whitespace
    // \W matches non-word characters, but we also want to keep spaces.
    // A cleaner way: replace everything that isn't a-z, 0-9, or whitespace with an empty string.
    // However, punctuation often sits between words (e.g., "times,it"), so we replace punctuation with spaces.
    const cleanedContent = lowerContent.replace(/[^a-z0-9\s]/g, ' ');

    // 3. Split into words and filter out empty strings
    const words = cleanedContent.split(/\s+/).filter(word => word.length > 0);

    // 4. Count frequencies
    const frequencyMap: Map<string, number> = new Map();
    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // 5. Convert Map to Array for sorting
    const sortedEntries = Array.from(frequencyMap.entries());

    // 6. Sort: count descending; for ties, sort alphabetically ascending
    sortedEntries.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA;
        }
        
        // Tie-breaker: alphabetical ascending
        return wordA.localeCompare(wordB);
    });

    // 7. Output
    for (const [word, count] of sortedEntries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();