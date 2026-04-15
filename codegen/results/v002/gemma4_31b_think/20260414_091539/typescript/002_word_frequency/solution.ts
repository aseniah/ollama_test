import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a text file, calculates the frequency of each word,
 * and prints them sorted by frequency (descending) and then alphabetically (ascending).
 */
function main(): void {
    const filePath = path.join('input', 'text.txt');
    
    let content: string;
    try {
        content = fs.readFileSync(filePath, 'utf8');
    } catch (error) {
        // Handle cases where the file might be missing or unreadable
        return;
    }

    // 1. Convert all words to lowercase
    // 2. Strip all punctuation / keep only letters
    // We use a regular expression that splits the string by any sequence of characters
    // that are NOT letters (a-z). This effectively removes punctuation and digits.
    const words = content
        .toLowerCase()
        .split(/[^a-z]+/)
        .filter(word => word.length > 0);

    // Count the frequency of each word using a Map
    const frequencyMap = new Map<string, number>();
    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // Convert the Map to an array of [word, count] tuples for sorting
    const sortedEntries = Array.from(frequencyMap.entries());

    // Sort by count descending; for ties, sort alphabetically ascending
    sortedEntries.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA; // Count descending
        }

        // Alphabetical ascending
        if (wordA < wordB) return -1;
        if (wordA > wordB) return 1;
        return 0;
    });

    // Output one line per unique word in the format: word: count
    for (const [word, count] of sortedEntries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

// Execute the program
main();