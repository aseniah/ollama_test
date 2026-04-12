import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = process.argv[2] || 'input/data.csv';

function parseCSV(content: string): any[][] {
  const lines = content.trim().split('\n');
  
  // Skip header row, return remaining as array of arrays
  return lines.slice(1).map(row => row.split(',').trim()) as any[];
}

// Read the CSV file
const csvContent = readFileSync(filePath, 'utf-8');

// Parse CSV into rows
const rows = parseCSV(csvContent);

// Convert to objects with typed fields
const data: any = [];
for (const row of rows) {
  data.push({
    Name: row[0],
    Age: parseInt(row[1]),
    Email: row[2],
    Score: parseFloat(row[3])
  });
}

// Output JSON to stdout
console.log(JSON.stringify(data, null, 2));