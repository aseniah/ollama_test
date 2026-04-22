import * as fs from 'fs';
import * as path from 'path';

function countWordFrequency(filePath: string): void {
    const content = fs.readFileSync(filePath, 'utf-8');
    const words = content.toLowerCase().split(/\s+/);
    
    const frequency: Record<string, number> = {};
    
    for (const rawWord of words) {
        // Strip punctuation: keep only letters
        const cleanWord = rawWord.replace(/[^a-z]/g, '');
        if (cleanWord.length === 0) {
            continue;
        }
        if (!frequency[cleanWord]) {
            frequency[cleanWord] = 0;
        }
        frequency[cleanWord]++;
    }
    
    // Convert to array for sorting
    const entries = Object.entries(frequency);
    
    // Sort by count descending, then alphabetically ascending for ties
    entries.sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });
    
    // Output
    for (const [word, count] of entries) {
        console.log(`${word}: ${count}`);
    }
}

const inputPath = path.join('input', 'text.txt');
countWordFrequency(inputPath);