import * as fs from 'fs';

const filePath = 'input/text.txt';
const content = fs.readFileSync(filePath, 'utf-8');

const words = content.toLowerCase().match(/[a-z]+/g) || [];

const wordCount: { [word: string]: number } = {};

for (const word of words) {
    wordCount[word] = (wordCount[word] || 0) + 1;
}

const sorted = Object.entries(wordCount).sort((a, b) => {
    if (b[1] !== a[1]) {
        return b[1] - a[1];
    }
    return a[0].localeCompare(b[0]);
});

for (const [word, count] of sorted) {
    console.log(`${word}: ${count}`);
}