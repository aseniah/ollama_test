import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = join('input', 'data.csv');
const fileContent = readFileSync(filePath, 'utf-8');

const lines = fileContent.trim().split('\n');
const headers = lines[0].split(',');

const result = lines.slice(1).map(line => {
  const values = line.split(',');
  return {
    Name: values[0].trim(),
    Age: parseInt(values[1].trim(), 10),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim())
  };
});

console.log(JSON.stringify(result));