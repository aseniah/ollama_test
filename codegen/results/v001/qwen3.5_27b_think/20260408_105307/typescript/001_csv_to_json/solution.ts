import * as fs from 'fs';
import * as path from 'path';

const filePath = 'input/data.csv';
const csvContent = fs.readFileSync(filePath, 'utf-8');

const lines = csvContent.split('\n').filter(line => line.trim() !== '');

const results: Array<{
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}> = [];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  const values = line.split(',');
  
  if (values.length >= 4) {
    results.push({
      Name: values[0].trim(),
      Age: parseInt(values[1].trim(), 10),
      Email: values[2].trim(),
      Score: parseFloat(values[3].trim())
    });
  }
}

console.log(JSON.stringify(results));