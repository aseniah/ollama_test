import * as fs from 'fs';
import * as path from 'path';

/**
 * The program reads a CSV file at 'input/data.csv' and converts it to a JSON array.
 * It parses specific columns: Name (string), Age (int), Email (string), and Score (float).
 */

function main() {
    const filePath = path.join('input', 'data.csv');

    try {
        // Read the file content as a UTF-8 string
        const content = fs.readFileSync(filePath, 'utf8');

        // Split content into lines, filtering out empty lines
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

        // If there are no lines or only a header, return an empty array
        if (lines.length <= 1) {
            console.log(JSON.stringify([]));
            return;
        }

        // Remove the header row
        const dataRows = lines.slice(1);

        const result = dataRows.map(row => {
            const columns = row.split(',');
            
            return {
                Name: columns[0]?.trim() || "",
                Age: parseInt(columns[1]?.trim() || "0", 10),
                Email: columns[2]?.trim() || "",
                Score: parseFloat(columns[3]?.trim() || "0")
            };
        });

        // Output as valid JSON to stdout
        process.stdout.write(JSON.stringify(result, null, 2) + '\n');

    } catch (error) {
        // Since the requirements state "do not output anything else" and output must be "valid JSON",
        // we typically wouldn't log errors unless necessary. 
        // However, for robustness in local dev, one might log to stderr.
        process.exit(1);
    }
}

main();