import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'data.csv');

// Read the CSV file
const csvContent = fs.readFileSync(inputPath, 'utf-8');

// Split lines and process rows
const lines = csvContent.trim().split('\n');

// Skip header line and process data rows
const header = lines[0]; // "Name,Age,Email,Score"
const fields = header.split(',');

const result: Array<{
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}> = [];

for (let i = 1; i < lines.length; i++) {
    const row = lines[i];
    const values = row.split(',');

    const name = values[0].trim();
    const age = parseInt(values[1].trim(), 10);
    const email = values[2].trim();
    const score = parseFloat(values[3].trim());

    result.push({
        Name: name,
        Age: age,
        Email: email,
        Score: score,
    });
}

// Output JSON array to stdout
console.log(JSON.stringify(result, null, 2));