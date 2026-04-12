import * as fs from 'fs';

const dataFilePath = 'input/data.csv';

// Read the file
const csvContent = fs.readFileSync(dataFilePath, 'utf-8');

// Split into lines
const lines = csvContent.trim().split('\n');

// Skip header line and parse data
const dataLines = lines.slice(1).map(line => line.split(',')).map(values => ({
    Name: values[0].trim(),
    Age: parseInt(values[1].trim()),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim())
}));

// Output JSON
console.log(JSON.stringify(dataLines, null, 2));