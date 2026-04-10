import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv' and transforms it 
 * into a JSON array of objects with specific types.
 */

function main() {
    const inputPath = path.join(process.cwd(), 'input', 'data.csv');

    try {
        // Check if file exists
        if (!fs.existsSync(inputPath)) {
            process.exit(1);
        }

        const content = fs.readFileSync(inputPath, 'utf-8');
        
        // Split by lines and remove empty lines
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

        if (lines.length < 2) {
            // Either empty file or only header
            process.stdout.write('[]');
            return;
        }

        // Extract header and rows
        // Assuming standard CSV format: Name,Age,Email,Score
        const rows = lines.slice(1);
        const result: Array<{
            Name: string;
            Age: number;
            Email: string;
            Score: number;
        }> = [];

        for (const row of rows) {
            const columns = row.split(',').map(col => col.trim());
            
            // Ensure we have enough columns to map to the required fields
            if (columns.length >= 4) {
                const name = columns[0];
                const age = parseInt(columns[1], 10);
                const email = columns[2];
                const score = parseFloat(columns[3]);

                // Only push if data is valid (Age and Score should be numbers)
                if (!isNaN(age) && !isNaN(score)) {
                    result.push({
                        Name: name,
                        Age: age,
                        Email: email,
                        Score: score
                    });
                }
            }
        }

        // Output the final JSON array to stdout
        process.stdout.write(JSON.stringify(result, null, 2));

    } catch (error) {
        // On error, we output nothing or exit, as per "Do not output anything else"
        process.exit(1);
    }
}

main();