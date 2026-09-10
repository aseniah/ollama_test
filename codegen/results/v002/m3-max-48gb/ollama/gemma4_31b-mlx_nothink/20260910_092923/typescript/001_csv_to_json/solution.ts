import * as fs from 'fs';
import * as path from 'path';

/**
 * The program reads a CSV file from input/data.csv,
 * parses it into an array of typed objects,
 * and outputs the result as a JSON string to stdout.
 */

interface UserRecord {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

function main() {
    const filePath = path.join('input', 'data.csv');

    try {
        // Read file content
        const content = fs.readFileSync(filePath, 'utf8');
        
        // Split by lines and remove empty lines at the end
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
        
        if (lines.length <= 1) {
            console.log(JSON.stringify([]));
            return;
        }

        // Extract header and data rows
        const header = lines[0].split(',');
        const dataRows = lines.slice(1);

        const result: UserRecord[] = dataRows.map(row => {
            const columns = row.split(',');
            
            return {
                Name: columns[0] ? columns[0].trim() : "",
                Age: parseInt(columns[1] || "0", 10),
                Email: columns[2] ? columns[2].trim() : "",
                Score: parseFloat(columns[3] || "0")
            };
        });

        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
    } catch (error) {
        // In a real scenario, we might log to stderr, 
        // but requirements state only valid JSON to stdout.
        process.exit(1);
    }
}

main();