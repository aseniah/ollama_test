import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv' and converts its content
 * into a JSON array of objects. Each object represents a row from the CSV,
 * with fields: Name (string), Age (integer), Email (string), and Score (float).
 */

function main() {
    const inputPath = path.join('input', 'data.csv');

    try {
        // Check if the file exists before attempting to read it.
        if (!fs.existsSync(inputPath)) {
            // If the file is missing, we output an empty array to ensure the output remains valid JSON.
            process.stdout.write('[]\n');
            return;
        }

        // Read the file content.
        const content = fs.readFileSync(inputPath, 'utf-8');

        // Split content into lines and remove empty lines (e.g., trailing newline).
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

        // If there's only a header or no content, return an empty array.
        if (lines.length <= 1) {
            process.stdout.write('[]\n');
            return;
        }

        // The first line is the header, so we skip it.
        const dataRows = lines.slice(1);

        // Transform each data row into an object.
        const result = dataRows.map(row => {
            // Split the row by comma and trim whitespace from each value.
            const columns = row.split(',').map(col => col.trim());

            // Destructure columns based on the expected schema.
            // Index 0: Name, Index 1: Age, Index 2: Email, Index 3: Score
            const name = columns[0];
            const age = parseInt(columns[1], 10);
            const email = columns[2];
            const score = parseFloat(columns[3]);

            return {
                Name: name,
                Age: isNaN(age) ? 0 : age,
                Email: email,
                Score: isNaN(score) ? 0.0 : score
            };
        });

        // Output the resulting array as a JSON string to stdout.
        // Using null and 2 for pretty-printing, but it can be JSON.stringify(result) for compact.
        process.stdout.write(JSON.stringify(result, null, 2) + '\n');

    } catch (error) {
        // In case of a critical error, we output an empty array to maintain valid JSON output constraint.
        process.stdout.write('[]\n');
    }
}

// Execute the script.
main();