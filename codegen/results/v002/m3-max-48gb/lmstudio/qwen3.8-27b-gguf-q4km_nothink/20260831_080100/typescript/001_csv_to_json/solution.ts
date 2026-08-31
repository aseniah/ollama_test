import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');
const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const row: { Name: string; Age: number; Email: string; Score: number } = {
    Name: values[0],
    Age: parseInt(values[1], 10),
    Email: values[2],
    Score: parseFloat(values[3]),
  };
  result.push(row);
}

console.log(JSON.stringify(result, null, 2));