import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'text.txt');
const content = fs.readFileSync(filePath, 'utf-8');

const wordCounts = new Map<string, number>();

const lines = content.split('\n');
for (const line of lines) {
    const words = line.split(/\s+/);
    for (const word of words) {
        const cleanedWord = word.toLowerCase().replace(/[^a-z]/g, '');
        if (cleanedWord.length > 0) {
            wordCounts.set(cleanedWord, (wordCounts.get(cleanedWord) || 0) + 1);
        }
    }
}

const sortedWords = Array.from(wordCounts.entries()).sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
});

for (const [word, count] of sortedWords) {
    console.log(`${word}: ${count}`);
}