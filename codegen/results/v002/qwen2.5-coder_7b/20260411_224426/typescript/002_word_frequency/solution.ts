import * as fs from 'fs';
import * as path from 'path';

// Helper function to check if a character is a letter
function isLetter(char: string): boolean {
    return /^[A-Za-z]$/.test(char);
}

// Read the file
const filePath = path.join(__dirname, 'input', 'text.txt');
fs.readFile(filePath, 'utf8', (err, data) => {
    if (err) {
        console.error('Error reading file:', err);
        return;
    }

    // Convert all words to lowercase and strip punctuation
    const words = data.toLowerCase().split(/\W+/);

    // Count the frequency of each word
    const wordCount: { [word: string]: number } = {};
    words.forEach(word => {
        if (word) {
            wordCount[word] = (wordCount[word] || 0) + 1;
        }
    });

    // Sort by count descending; for ties, sort alphabetically ascending
    const sortedWords = Object.entries(wordCount).sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        } else {
            return a[0].localeCompare(b[0]);
        }
    });

    // Output one line per unique word
    sortedWords.forEach(([word, count]) => {
        console.log(`${word}: ${count}`);
    });
});