import { readFileSync } from 'fs';
import path from 'path';

const filePath = path.resolve(process.cwd(), 'input/data.csv');
const content = readFileSync(filePath, 'utf-8');

const lines = content.trim().split('\n');
const header = lines[0].split(',');

const rows: { Name: string; Age: number; Email: string; Score: number }[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const row: { Name: string; Age: number; Email: string; Score: number } = {
    Name: values[0].trim(),
    Age: parseInt(values[1].trim(), 10),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim()),
  };
  rows.push(row);
}

process.stdout.write(JSON.stringify(rows, null, 2) + '\n');