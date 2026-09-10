import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a CSV file from 'input/data.csv', 
 * parses its contents based on the specified schema, 
 * and outputs a JSON array to stdout.
 */

function main() {
    const filePath = path.join('input', 'data.csv');

    try {
        // Read the content of the CSV file
        const data = fs.readFileSync(filePath, 'utf8');
        
        // Split into lines and filter out any trailing empty lines
        const lines = data.split(/\r?\n/).filter(line => line.trim() !== '');
        
        if (lines.length === 0) {
            console.log(JSON.stringify([]));
            return;
        }

        // The first line is the header, which we skip
        const rows = lines.slice(1);
        
        const result = rows.map(row => {
            const columns = row.split(',');
            
            // Extract fields based on the expected positions:
            // Name: index 0, Age: index 1, Email: index 2, Score: index 3
            return {
                Name: columns[0]?.trim() || "",
                Age: parseInt(columns[1]?.trim() || "0", 10),
                Email: columns[2]?.trim() || "",
                Score: parseFloat(columns[3]?.trim() || "0")
            };
        });

        // Output the final array as valid JSON
        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
    } catch (err) {
        // If there's an error reading the file, we exit quietly or handle it
        // as per requirements (output must be valid JSON or nothing else).
        process.exit(1);
    }
}

main();