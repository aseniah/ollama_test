import * as fs from 'fs';
import * as path from 'path';

/**
 * The program reads the file `input/text.txt`, cleans the text by keeping only letters,
 * counts the frequency of each lowercase word, and outputs them sorted by frequency
 * (descending) and then alphabetically (ascending).
 */
function solve() {
    const filePath = path.join('input', 'text.txt');
    
    let text: string;
    try {
        text = fs.readFileSync(filePath, 'utf8');
    } catch (err) {
        // If file doesn't exist or cannot be read, exit silently or handle as needed.
        // Given the prompt, we assume the file exists.
        return;
    }

    // 1. Convert to lowercase and find all sequences of letters (stripping punctuation/numbers)
    const words = text.toLowerCase().match(/[a-z]+/g) || [];

    // 2. Count frequencies using a Map
    const frequencyMap = new Map<string, number>();
    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // 3. Convert Map to an array for sorting
    const sortedWords = Array.from(frequencyMap.entries());

    // 4. Sort by count descending, then alphabetically ascending
    sortedWords.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA; // Descending frequency
        }
        return wordA.localeCompare(wordB); // Ascending alphabetical
    });

    // 5. Output result in "word: count" format
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();