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

    // Convert to lowercase and split by whitespace
    const rawWords = content.toLowerCase().split(/\s+/);
    const frequencyMap: Map<string, number> = new Map();

    for (const rawWord of rawWords) {
        // Strip all punctuation: keep only letters (a-z)
        const cleanWord = rawWord.replace(/[^a-z]/g, '');
        
        if (cleanWord.length > 0) {
            frequencyMap.set(cleanWord, (frequencyMap.get(cleanWord) || 0) + 1);
        }
    }

    // Convert map to array for sorting
    const sortedWords = Array.from(frequencyMap.entries());

    sortedWords.sort((a, b) => {
        const [wordA, countA] = a;
        const [wordB, countB] = b;

        // Primary sort: count descending
        if (countB !== countA) {
            return countB - countA;
        }
        // Secondary sort: alphabetical ascending
        return wordA.localeCompare(wordB);
    });

    // Output results
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();