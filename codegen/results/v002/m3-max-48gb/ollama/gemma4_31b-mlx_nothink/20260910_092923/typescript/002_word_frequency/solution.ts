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

    // Convert to lowercase
    const lowerContent = content.toLowerCase();
    
    // Split by whitespace and clean punctuation
    // regex [^a-z] removes anything that is not a lowercase letter
    const words = lowerContent.split(/\s+/);
    const wordCounts: Map<string, number> = new Map();

    for (const rawWord of words) {
        const cleanWord = rawWord.replace(/[^a-z]/g, '');
        if (cleanWord.length > 0) {
            wordCounts.set(cleanWord, (wordCounts.get(cleanWord) || 0) + 1);
        }
    }

    // Convert Map to Array of tuples [word, count]
    const sortedWords = Array.from(wordCounts.entries());

    // Sort: count descending, then word ascending
    sortedWords.sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });

    // Output results
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();