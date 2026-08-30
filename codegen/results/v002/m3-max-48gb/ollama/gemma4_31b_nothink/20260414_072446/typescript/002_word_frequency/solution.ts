import * as fs from 'fs';
import * as path from 'path';

function solve() {
    const filePath = path.join('input', 'text.txt');
    
    let content: string;
    try {
        content = fs.readFileSync(filePath, 'utf8');
    } catch (err) {
        return;
    }

    // Convert to lowercase and replace any non-letter character with a space
    // This effectively strips punctuation and splits words
    const cleanedText = content.toLowerCase().replace(/[^a-z\s]/g, ' ');
    const words = cleanedText.split(/\s+/).filter(word => word.length > 0);

    const freqMap: Map<string, number> = new Map();

    for (const word of words) {
        freqMap.set(word, (freqMap.get(word) || 0) + 1);
    }

    // Convert map to array for sorting
    const sortedWords = Array.from(freqMap.entries()).sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        if (countB !== countA) {
            return countB - countA; // Sort by count descending
        }
        return wordA.localeCompare(wordB); // Sort alphabetically ascending
    });

    // Output in the format word: count
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();