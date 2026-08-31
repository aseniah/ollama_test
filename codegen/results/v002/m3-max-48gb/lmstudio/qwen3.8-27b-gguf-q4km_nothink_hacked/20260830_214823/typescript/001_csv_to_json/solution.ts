import * as fs from 'fs';

// Read the CSV file
const csvContent = fs.readFileSync('input/data.csv', 'utf-8');

// Split into lines
const lines = csvContent.trim().split('\n');

// Skip header
const header = lines[0].split(',');
const rows = lines.slice(1);

// Map header to field names
const result = rows.map(row => {
    const parts = row.split(',');
    return {
        Name: parts[0].trim(),
        Age: parseInt(parts[1].trim(), 10),
        Email: parts[2].trim(),
        Score: parseFloat(parts[3].trim())
    };
});

// Output JSON
console.log(JSON.stringify(result, null, 2));