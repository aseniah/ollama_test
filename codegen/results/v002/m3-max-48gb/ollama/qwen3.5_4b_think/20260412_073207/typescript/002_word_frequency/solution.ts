import * as fs from 'fs';

// Read the file content
const content = fs.readFileSync('input/text.txt', 'utf8');

// Convert to lowercase and split into words (keeping only letters)
const words = content.toLowerCase()
    .replace(/[^a-z\s]/g, ' ')
    .trim()
    .split(/\s+/);

// Count frequencies
const wordCount: Record<string, number> = {};
for (const word of words) {
    const trimmedWord = word;
    if (trimmedWord.length > 0) {
        wordCount[trimmedWord] = (wordCount[trimmedWord] || 0) + 1;
    }
}

// Create array of [word, count]
const sortedWords = Object.entries(wordCount).map(([word, count]) => ({ word, count }));

// Sort by count descending, then alphabetically ascending
sortedWords.sort((a, b) => {
    if (b.count !== a.count) {
        return b.count - a.count;
    }
    return a.word.localeCompare(b.word);
});

// Output
for (const { word, count } of sortedWords) {
    console.log(`${word}: ${count}`);
}