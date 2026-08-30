import fs from 'fs';
import path from 'path';

// Function to process a single line
function processLine(line: string): string[] {
    const words = line.toLowerCase().replace(/[^\w\s]/g, '').split(/\s+/);
    return words;
}

// Main function to read and process the file
async function main() {
    const inputPath = path.join(__dirname, 'input', 'text.txt');
    const outputPath = path.join(__dirname, 'output', 'word_frequencies.txt');

    try {
        const lines = await fs.readdir(inputPath);
        const text = lines.find(file => file === 'text.txt').fileContents;

        const wordCounts: Record<string, number> = {};

        text.split('\n').forEach(line => {
            const words = processLine(line);
            words.forEach(word => {
                if (wordCounts.has(word)) {
                    wordCounts[word]++;
                } else {
                    wordCounts[word] = 1;
                }
            });
        });

        // Sort the word counts by count descending, then alphabetically ascending
        const sortedWords = Object.entries(wordCounts)
            .sort((a, b) => {
                const countComparison = b[1] - a[1];
                if (countComparison !== 0) {
                    return countComparison;
                }
                return a[0].localeCompare(b[0]);
            })
            .map(word => `${word[0]}: ${word[1]}`);

        // Write sorted words to output file
        fs.writeFileSync(outputPath, sortedWords.join('\n'));

        console.log(`Word frequencies have been written to output/word_frequencies.txt`);
    } catch (error) {
        console.error('An error occurred:', error);
    }
}

main();