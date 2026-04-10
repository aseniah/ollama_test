import * as fs from 'fs';

// Read the CSV file
const content = fs.readFileSync('input/data.csv', 'utf-8');

// Split into lines
const lines = content.trim().split('\n');

// Parse data rows (skip header)
const result: Array<{ name: string; age: number; email: string; score: number }> = [];

// Remove the header row
const dataRows = lines.slice(1);

// Parse each data row
for (const row of dataRows) {
  const fields = row.split(',').map(f => f.trim());
  
  if (fields.length < 4) continue;
  
  result.push({
    name: fields[0],
    age: parseInt(fields[1], 10),
    email: fields[2],
    score: parseFloat(fields[3])
  });
}

// Output as JSON
console.log(JSON.stringify(result));