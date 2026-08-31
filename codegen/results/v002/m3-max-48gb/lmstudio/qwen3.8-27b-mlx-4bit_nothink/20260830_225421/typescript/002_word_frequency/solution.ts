import * as fs from 'fs';
import * as path from 'path';

// Define the input file path
const inputPath = path.join('input', 'text.txt');

try {
    // 1. Read the file
    const data = fs.readFileSync(inputPath, 'utf8');
    
    // 2. Process the text
    // Convert to lowercase
    let processed = data.toLowerCase();
    
    // Strip all punctuation, keep only letters.
    // Replace any character that is not a letter with a space to preserve word boundaries.
    processed = processed.replace(/[^a-z]/g, ' ');
    
    // Split by whitespace (spaces, newlines, etc.)
    const words = processed.split(/\s+/).filter(word => word.length > 0);
    
    // 3. Count frequencies
    const freq: Record<string, number> = {};
    for (const word of words) {
        if (!freq[word]) {
            freq[word] = 0;
        }
        freq[word]++;
    }
    
    // 4. Prepare for sorting
    // Convert object to array of [word, count] pairs
    const entries = Object.entries(freq);
    
    // Sort logic:
    // 1. Count descending
    // 2. Word ascending (alphabetical)
    entries.sort((a, b) => {
        const wordA = a[0];
        const countA = a[1];
        const wordB = b[0];
        const countB = b[1];
        
        if (countB !== countA) {
            return countB - countA; // Descending count
        }
        
        // If counts are equal, sort alphabetically ascending
        return wordA.localeCompare(wordB);
    });
    
    // 5. Output
    for (const [word, count] of entries) {
        console.log(`${word}: ${count}`);
    }
} catch (error) {
    console.error(`Error reading file: ${error}`);
}