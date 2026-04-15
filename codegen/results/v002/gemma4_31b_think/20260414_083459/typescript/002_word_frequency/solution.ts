import * as fs from 'fs';
import * as path from 'path';

/**
 * Solves the word frequency counting task.
 */
function solve() {
    const filePath = path.join('input', 'text.txt');
    
    let content: string;
    try {
        content = fs.readFileSync(filePath, 'utf8');
    } catch (error) {
        // The prompt implies the file exists, but we'll handle errors gracefully.
        return;
    }

    // Convert to lowercase
    const lowerContent = content.toLowerCase();

    // Extract words containing only letters
    // This strips all punctuation and keeps only alphabetical sequences
    const words = lowerContent.match(/[a-z]+/g);

    if (!words) {
        return;
    }

    // Count frequencies
    const frequencyMap = new Map<string, number>();
    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // Convert Map to array for sorting
    const sortedWords = Array.from(frequencyMap.entries());

    // Sort logic:
    // 1. Count descending
    // 2. Alphabetical ascending (for ties)
    sortedWords.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA;
        }
        return wordA.localeCompare(wordB);
    });

    // Output formatted results
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

// Run the solver
solve();