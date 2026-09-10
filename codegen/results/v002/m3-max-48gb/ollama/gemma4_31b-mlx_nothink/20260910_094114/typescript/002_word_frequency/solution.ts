import * as fs from 'fs';
import * as path from 'path';

function solve() {
    const filePath = path.join('input', 'text.txt');
    
    let content: string;
    try {
        content = fs.readFileSync(filePath, 'utf8');
    } catch (err) {
        return;
    }

    // Rule: Convert to lowercase
    const lowerContent = content.toLowerCase();
    
    // Rule: Strip all punctuation (keep only letters and spaces to split words)
    // This regex replaces anything that isn't a lowercase letter or whitespace with an empty string
    const cleanedContent = lowerContent.replace(/[^a-z\s]/g, ' ');
    
    // Split into words by whitespace and filter out empty strings
    const words = cleanedContent.split(/\s+/).filter(word => word.length > 0);
    
    const frequencyMap: Map<string, number> = new Map();
    
    for (const word of words) {
        frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
    }
    
    // Convert map to array for sorting
    const sortedWords = Array.from(frequencyMap.entries()).sort(([wordA, countA], [wordB, countB]) => {
        // Rule: Sort by count descending
        if (countB !== countA) {
            return countB - countA;
        }
        // Rule: Ties sort alphabetically ascending
        return wordA.localeCompare(wordB);
    });
    
    // Rule: Output one line per unique word in the format: word: count
    for (const [word, count] of sortedWords) {
        process.stdout.write(`${word}: ${count}\n`);
    }
}

solve();