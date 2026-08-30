import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const filePath = path.join(process.cwd(), 'input/data.csv');
const data = fs.readFileSync(filePath, 'utf-8');

// Split by newline to get rows
const lines = data.trim().split('\n');

// Skip header row (first line), process remaining rows
const result = lines.slice(1).map(line => {
  const parts = line.split(',');
  
  return {
    Name: parts[0].trim(),
    Age: parseInt(parts[1].trim(), 10),
    Email: parts[2].trim(),
    Score: parseFloat(parts[3].trim())
  };
});

// Output JSON array to stdout
console.log(JSON.stringify(result));