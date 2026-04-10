import * as fs from 'fs';
import * as path from 'path';

// Function to read file and count word frequencies
function countWordFrequencies(filePath: string): void {
    try {
        // Read the file content
        const fileContent = fs.readFileSync(filePath, 'utf8');

        // Convert text to lowercase
        const lowercasedText = fileContent.toLowerCase();

        // Remove punctuation (keep only letters)
        const cleanedText = lowercasedText.replace(/[^\w\s]/g, '');

        // Split the text into words
        const words = cleanedText.split(/\s+/);

        // Create a map to store word frequencies
        const frequencyMap: { [word: string]: number } = {};

        // Count word frequencies
        words.forEach(word => {
            if (frequencyMap[word]) {
                frequencyMap[word]++;
            } else {
                frequencyMap[word] = 1;
            }
        });

        // Convert map to an array of [word, count] pairs
        const freqArray: [string, number][] = Object.entries(frequencyMap);

        // Sort by count descending; for ties, sort alphabetically ascending
        freqArray.sort((a, b) => {
            if (b[1] === a[1]) {
                return a[0].localeCompare(b[0]);
            }
            return b[1] - a[1];
        });

        // Output one line per unique word in the format: `word: count`
        freqArray.forEach(([word, count]) => {
            console.log(`${word}: ${count}`);
        });
    } catch (error) {
        console.error(`Error reading file: ${error.message}`);
    }
}

// Get the file path from command-line argument
const filePath = path.join(__dirname, 'input', 'text.txt');

// Call the function with the provided file path
countWordFrequencies(filePath);