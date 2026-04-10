import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header line
const dataRows = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const line of dataRows) {
  // Split by comma, handling potential basic cases. 
  // Note: This assumes CSV does not contain commas within quoted fields.
  const parts = line.split(',');
  
  if (parts.length >= 4) {
    result.push({
      Name: parts[0],
      Age: parseInt(parts[1], 10),
      Email: parts[2],
      Score: parseFloat(parts[3])
    });
  }
}

console.log(JSON.stringify(result));