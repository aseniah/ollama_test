import { readFileSync } from 'fs';
import path from 'path';
import { Tokenizer, Punctuation } from '@types/tokenizer';

// Helper function to clean words
function cleanWord(word: string): string {
    const cleanedWord = word.toLowerCase();
    // Remove punctuation
    cleanedWord = cleanedWord.replace(/[^\w\s]/g, '');
    return cleanedWord;
}

async function main() {
    const filePath = path.join(__dirname, 'input', 'text.txt');
    const content = await readFileSync(filePath, 'utf8');

    const words = Tokenizer.tokenize(content, Tokenizer.Alphabetic, Tokenizer.Word);
    
    const wordCount: Record<string, number> = {};

    words.forEach(word => {
        const cleanedWord = cleanWord(word);
        if (cleanedWord) {
            wordCount[cleanedWord] = (wordCount[cleanedWord] || 0) + 1;
        }
    });

    const sortedWords = Object.entries(wordCount)
        .sort((a, b) => {
            if (a[1] !== b[1]) {
                return b[1] - a[1]; // Sort by count descending
            }
            return a[0].localeCompare(b[0]); // Sort alphabetically ascending for ties
        })
        .map(([word, count]) => `${word}: ${count}`);

    console.log(sortedWords.join('\n'));
}

main();