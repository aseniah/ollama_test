import * as fs from 'fs';
import * as path from 'path';

/**
 * Word frequency counter
 * Reads input/text.txt, cleans words, and sorts them based on frequency and alphabet.
 */
function main() {
    const inputPath = path.join(process.cwd(), 'input', 'text.txt');

    let content: string;
    try {
        content = fs.readFileSync(inputPath, 'utf8');
    } catch (err) {
        return;
    }

    // Convert to lowercase and replace non-letter characters with spaces
    // This handles punctuation and numbers as requested (keep only letters)
    const cleanedText = content.toLowerCase().replace(/[^a-z\s]/g, ' ');
    
    // Split by whitespace and filter out empty strings
    const words = cleanedText.split(/\s+/).filter(word => word.length > 0);

    const frequencyMap: Map<string, number> = new Map();

    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // Convert map to array of [word, count] pairs
    const sortedWords = Array.from(frequencyMap.entries());

    // Sorting logic:
    // 1. Frequency descending (b[1] - a[1])
    // 2. Alphabetical ascending (a[0].localeCompare(b[0]))
    sortedWords.sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });

    // Output format: "word: count"
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

main();