import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header (first line)
const dataLines = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const line of dataLines) {
  if (!line.trim()) continue; // Skip empty lines
  
  const values = line.split(',');
  
  result.push({
    Name: values[0]?.trim() || '',
    Age: parseInt(values[1], 10),
    Email: values[2]?.trim() || '',
    Score: parseFloat(values[3])
  });
}

console.log(JSON.stringify(result));