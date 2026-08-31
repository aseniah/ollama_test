import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'data.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');

const lines = csvContent.split('\n').filter(line => line.trim() !== '');
const header = lines[0].split(',');
const rows: object[] = [];

for (let i = 1; i < lines.length; i++) {
  const fields = lines[i].split(',');
  const obj: { Name: string; Age: number; Email: string; Score: number } = {
    Name: fields[0],
    Age: parseInt(fields[1], 10),
    Email: fields[2],
    Score: parseFloat(fields[3])
  };
  rows.push(obj);
}

console.log(JSON.stringify(rows, null, 2));