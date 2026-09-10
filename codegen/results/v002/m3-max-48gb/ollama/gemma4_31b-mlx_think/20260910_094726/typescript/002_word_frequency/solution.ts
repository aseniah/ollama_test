import * as fs from 'fs';
import * as path from 'path';

function solve() {
    const filePath = path.join('input', 'text.txt');
    
    try {
        // Read the file content
        const content = fs.readFileSync(filePath, 'utf8');
        
        // Convert to lowercase and extract words (keeping only letters)
        // Using a regex to find all sequences of alphabetic characters
        const words = content.toLowerCase().match(/[a-z]+/g);
        
        if (!words) return;

        // Count frequencies
        const frequencyMap = new Map<string, number>();
        for (const word of words) {
            frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
        }

        // Convert map to array for sorting
        const sortedEntries = Array.from(frequencyMap.entries()).sort((a, b) => {
            const [wordA, countA] = a;
            const [wordB, countB] = b;

            // Primary sort: count descending
            if (countB !== countA) {
                return countB - countA;
            }
            // Secondary sort: alphabetically ascending
            return wordA.localeCompare(wordB);
        });

        // Output results
        for (const [word, count] of sortedEntries) {
            process.stdout.write(`${word}: ${count}\n`);
        }
    } catch (err) {
        // Silent error handling as per "do not output anything else" 
        // unless it's the required format.
    }
}

solve();