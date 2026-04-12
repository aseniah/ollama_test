import * as fs from 'fs';
import * as path from 'path';

const csvPath = 'input/data.csv';

// Read the CSV file
const csvContent = fs.readFileSync(csvPath, 'utf8');

// Split into lines and get header
const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');

const data = [];

for (let i = 1; i < lines.length; i++) {
  if (!lines[i].trim()) continue; // skip empty rows
  
  const values = lines[i].split(',');
  
  const row: Record<string, any> = {};
  row.Name = values[0]?.trim() as string;
  row.Age = parseInt(values[1]);
  row.Email = values[2]?.trim() as string;
  row.Score = parseFloat(values[3]);
  
  data.push(row);
}

console.log(JSON.stringify(data));