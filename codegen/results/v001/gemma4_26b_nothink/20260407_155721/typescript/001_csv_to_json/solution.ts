import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv', parses it,
 * and outputs a JSON array to stdout.
 */

function main() {
    const filePath = path.join(process.cwd(), 'input', 'data.csv');

    try {
        // Check if file exists
        if (!fs.existsSync(filePath)) {
            process.stderr.write(`Error: File not found at ${filePath}\n`);
            process.exit(1);
        }

        const content = fs.readFileSync(filePath, 'utf-8');
        const lines = content.split(/\r?\n/);

        // Filter out empty lines
        const rows = lines.filter(line => line.trim() !== '');

        if (rows.length === 0) {
            console.log('[]');
            return;
        }

        // Assume first row is the header
        const header = rows[0].split(',').map(h => h.trim());
        
        // Map headers to indices for robust parsing
        const nameIdx = header.indexOf('Name');
        const ageIdx = header.indexOf('Age');
        const emailIdx = header.indexOf('Email');
        const scoreIdx = header.indexOf('Score');

        // Validate header existence
        if (nameIdx === -1 || ageIdx === -1 || emailIdx === -1 || scoreIdx === -1) {
            process.stderr.write('Error: CSV header must contain Name, Age, Email, and Score\n');
            process.exit(1);
        }

        const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

        // Iterate through data rows (skipping header)
        for (let i = 1; i < rows.length; i++) {
            const values = rows[i].split(',').map(v => v.trim());
            
            // Ensure row has enough columns
            if (values.length <= Math.max(nameIdx, ageIdx, emailIdx, scoreIdx)) {
                continue; 
            }

            const obj = {
                Name: values[nameIdx],
                Age: parseInt(values[ageIdx], 10),
                Email: values[emailIdx],
                Score: parseFloat(values[scoreIdx])
            };

            // Check if parsing failed for numbers
            if (isNaN(obj.Age) || isNaN(obj.Score)) {
                // Skip malformed rows or handle error
                continue;
            }

            result.push(obj);
        }

        // Output valid JSON to stdout
        process.stdout.write(JSON.stringify(result, null, 2) + '\n');

    } catch (error) {
        if (error instanceof Error) {
            process.stderr.write(`Error processing file: ${error.message}\n`);
        } else {
            process.stderr.write('An unknown error occurred\n');
        }
        process.exit(1);
    }
}

main();