import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a CSV file and converts it to a JSON array
 * based on specific field types.
 */

function solve() {
    const filePath = path.join('input', 'data.csv');
    
    try {
        // Read the file content synchronously
        const content = fs.readFileSync(filePath, 'utf8');
        
        // Split by line and filter out empty lines
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
        
        if (lines.length === 0) {
            console.log(JSON.stringify([]));
            return;
        }

        // The first line is the header, we slice it off
        const dataLines = lines.slice(1);
        
        const result = dataLines.map(line => {
            const columns = line.split(',');
            
            // Expected format: Name, Age, Email, Score
            return {
                Name: columns[0] ? columns[0].trim() : "",
                Age: columns[1] ? parseInt(columns[1].trim(), 10) : 0,
                Email: columns[2] ? columns[2].trim() : "",
                Score: columns[3] ? parseFloat(columns[3].trim()) : 0.0
            };
        });

        // Output the final array as a JSON string
        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
        
    } catch (error) {
        // In a real scenario, we might handle errors differently, 
        // but based on the prompt, we just need to ensure valid JSON output.
        process.exit(1);
    }
}

solve();