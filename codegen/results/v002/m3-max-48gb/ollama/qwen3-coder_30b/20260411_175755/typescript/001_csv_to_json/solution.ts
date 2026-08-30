import { readFileSync } from 'fs';
import { join } from 'path';

// Read the CSV file
const filePath = join('input', 'data.csv');
const csvContent = readFileSync(filePath, 'utf-8');

// Split into lines
const lines = csvContent.trim().split('\n');

// Get header row
const headers = lines[0].split(',');

// Process each data row
const result = lines.slice(1).map(line => {
  const values = line.split(',');
  return {
    Name: values[0].trim(),
    Age: parseInt(values[1].trim(), 10),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim())
  };
});

// Output as JSON
console.log(JSON.stringify(result));