import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const csvPath = path.join('input', 'data.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

// Split into lines
const lines = csvContent.trim().split('\n');

// Skip header, parse each row
const results = lines.slice(1).map(line => {
    const parts = line.split(',');
    return {
        Name: parts[0].trim(),
        Age: parseInt(parts[1].trim(), 10),
        Email: parts[2].trim(),
        Score: parseFloat(parts[3].trim())
    };
});

// Output as JSON
console.log(JSON.stringify(results));