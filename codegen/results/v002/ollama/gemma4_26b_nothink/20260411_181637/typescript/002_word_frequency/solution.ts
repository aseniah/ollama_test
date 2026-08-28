import * as fs from 'fs';
import * as path from 'path';

/**
 * Solves the task of counting word frequencies from a specific file.
 * Rules:
 * - Convert to lowercase
 * - Strip punctuation (keep only letters)
 * - Output format: "word: count"
 * - Sort: Count (descending), then Word (alphabetical ascending)
 */

function solve() {
    const filePath = path.join(process.cwd(), 'input', 'text.txt');

    let content: string;

    try {
        content = fs.readFileSync(filePath, 'utf-8');
    } catch (err) {
        // If file doesn't exist, we exit silently as per "Do not output anything else"
        // but in a real environment, we'd handle this error.
        return;
    }

    // 1. Convert to lowercase
    const lowercased = content.toLowerCase();

    // 2. Strip all punctuation (keep only letters and whitespace for splitting)
    // We replace everything that is not a lowercase letter or whitespace with an empty string.
    // We also replace newlines with spaces to ensure words aren't merged.
    const cleaned = lowercased.replace(/[^a-z\s]/g, ' ').replace(/\s+/g, ' ');

    // 3. Split into words and filter out empty strings
    const words = cleaned.trim().split(' ').filter(word => word.length > 0);

    // 4. Count frequencies
    const frequencyMap: Record<string, number> = {};
    for (const word of words) {
        frequencyMap[word] = (frequencyMap[word] || 0) + 1;
    }

    // 5. Convert map to array for sorting
    const sortedEntries = Object.entries(frequencyMap).sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        const countDiff = countB - countA; // Count descending
        if (countDiff !== 0) {
            return countDiff;
        }

        // For ties, sort alphabetically ascending
        return wordA.localeCompare(wordable);
    });

    // Note: Re-implementing sort logic clearly
    const entries = Object.entries(frequencyMap);
    entries.sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });

    // 6. Output results
    for (const [word, count] of entries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();