import fs from 'fs';
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const filePath = join(__dirname, 'input', 'text.txt');

const content = fs.readFileSync(filePath, 'utf-8');

const words = content.toLowerCase().split(/\s+/);

const wordCount = new Map();
for (const word of words) {
    const cleanWord = word.replace(/[^a-z]/g, '');
    if (cleanWord.length > 0) {
        const existing = wordCount.get(cleanWord);
        wordCount.set(cleanWord, existing + 1);
    }
}

const wordsArray = Array.from(wordCount.entries());
wordsArray.sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
});

for (const [word, count] of wordsArray) {
    console.log(`${word}: ${count}`);
}