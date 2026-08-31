import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.resolve('input/text.txt');

try {
    const content = fs.readFileSync(inputPath, 'utf-8');
    
    // Convert to lowercase
    const lower = content.toLowerCase();
    
    // Strip all non-letters (keep only a-z)
    const lettersOnly = lower.replace(/[^a-z\s]/g, ' ');
    
    // Split by whitespace and filter empty strings
    const words = lettersOnly.split(/\s+/).filter(w => w.length > 0);
    
    // Count frequencies
    const freq: Map<string, number> = new Map();
    for (const word of words) {
        freq.set(word, (freq.get(word) || 0) + 1);
    }
    
    // Convert to array of [word, count] pairs
    const entries = Array.from(freq.entries());
    
    // Sort by count descending; for ties, sort alphabetically ascending
    entries.sort((a, b) => {
        if (b[1] !== a[1]) {
            return b[1] - a[1];
        }
        return a[0].localeCompare(b[0]);
    });
    
    // Output each word: count
    const output = entries.map(([word, count]) => `${word}: ${count}`).join('\n');
    
    console.log(output);
    
} catch (error) {
    // If file doesn't exist, output nothing
    // The problem says "Do not output anything else" so we just don't output anything
}