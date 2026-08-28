import * as fs from 'fs';
import * as path from 'path';

function main() {
    const filePath = process.argv[2] || 'input/text.txt';
    
    try {
        const text = fs.readFileSync(filePath, 'utf-8');
        
        // Convert to lowercase and keep only letters
        const normalized = text.toLowerCase().replace(/[^a-z\s]/g, ' ');
        
        // Split into words and filter out empty strings
        const words = normalized.split(/\s+/).filter(word => word.length > 0);
        
        // Count frequencies
        const frequencyMap = new Map<string, number>();
        for (const word of words) {
            frequencyMap.set(word, (frequencyMap.get(word) || 0) + 1);
        }
        
        // Convert to array and sort
        const sorted = Array.from(frequencyMap.entries()).sort((a, b) => {
            if (b[1] !== a[1]) {
                return b[1] - a[1]; // Descending by count
            }
            return a[0].localeCompare(b[0]); // Ascending alphabetically for ties
        });
        
        // Output results
        for (const [word, count] of sorted) {
            console.log(`${word}: ${count}`);
        }
    } catch (error) {
        console.error(`Error reading file: ${(error as Error).message}`);
        process.exit(1);
    }
}

main();