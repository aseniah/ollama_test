import * as fs from 'fs';
import * as path from 'path';

interface UserRecord {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

function main() {
    const filePath = path.join('input', 'data.csv');

    try {
        const content = fs.readFileSync(filePath, 'utf8');
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

        if (lines.length === 0) {
            process.stdout.write(JSON.stringify([], null, 2));
            return;
        }

        // Remove the header row
        const dataRows = lines.slice(1);
        const results: UserRecord[] = [];

        for (const row of dataRows) {
            const columns = row.split(',');

            if (columns.length >= 4) {
                results.push({
                    Name: columns[0].trim(),
                    Age: parseInt(columns[1].trim(), 10),
                    Email: columns[2].trim(),
                    Score: parseFloat(columns[3].trim())
                });
            }
        }

        process.stdout.write(JSON.stringify(results, null, 2));
    } catch (error) {
        // Since the requirement is to output valid JSON and nothing else, 
        // we output an empty array or a JSON error if the file is missing.
        process.stdout.write(JSON.stringify({ error: "Could not read file" }));
        process.exit(1);
    }
}

main();