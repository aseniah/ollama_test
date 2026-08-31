import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const csvContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines
const lines = csvContent.split('\n').filter(line => line.trim() !== '');

// Skip the header
const header = lines[0].split(',');
const rows = lines.slice(1);

// Parse each row into an object
const results: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const row of rows) {
    const fields = row.split(',');
    const name = fields[0].trim();
    const age = parseInt(fields[1].trim(), 10);
    const email = fields[2].trim();
    const score = parseFloat(fields[3].trim());
    
    results.push({
        Name: name,
        Age: age,
        Email: email,
        Score: score
    });
}

// Output as JSON
console.log(JSON.stringify(results, null, 2));