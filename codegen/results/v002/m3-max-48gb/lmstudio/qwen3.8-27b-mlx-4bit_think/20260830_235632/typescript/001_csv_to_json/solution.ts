import * as fs from 'fs';
import * as path from 'path';

const csvPath = path.resolve('input/data.csv');
const content = fs.readFileSync(csvPath, 'utf-8');

const lines = content.split('\n').filter(line => line.trim() !== '');
const header = lines[0].split(',').map(h => h.trim());

const rows: object[] = lines.slice(1).map(line => {
  const values = line.split(',').map(v => v.trim());
  const obj: { Name: string; Age: number; Email: string; Score: number } = {
    Name: values[header.indexOf('Name')] ?? '',
    Age: parseInt(values[header.indexOf('Age')] ?? '0', 10),
    Email: values[header.indexOf('Email')] ?? '',
    Score: parseFloat(values[header.indexOf('Score')] ?? '0'),
  };
  return obj;
});

process.stdout.write(JSON.stringify(rows, null, 2) + '\n');