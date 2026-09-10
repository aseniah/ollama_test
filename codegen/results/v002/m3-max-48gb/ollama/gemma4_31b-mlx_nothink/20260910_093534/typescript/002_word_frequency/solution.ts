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

    // Convert to lowercase and replace anything that isn't a letter or whitespace with an empty string
    const cleanedText = content.toLowerCase().replace(/[^a-z\s]/g, ' ');
    
    // Split by whitespace and filter out empty strings
    const words = cleanedText.split(/\s+/).filter(word => word.length > 0);

    const frequencyMap: Map<string, number> = new Map();

    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }

    // Convert map to array for sorting
    const sortedWords = Array.from(frequencyMap.entries()).sort(([wordA, countA], [wordB, countB]) => {
        if (countB !== countA) {
            return countB - countA; // Sort by count descending
        }
        return wordA.localeCompare(wordB); // Sort alphabetically ascending for ties
    });

    // Output result
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();