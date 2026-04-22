import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(__dirname, 'input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');
const header = lines[0].split(',');

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj: any = {
    Name: values[0],
    Age: parseInt(values[1], 10),
    Email: values[2],
    Score: parseFloat(values[3]),
  };
  result.push(obj);
}

console.log(JSON.stringify(result));