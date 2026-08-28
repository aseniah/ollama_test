import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = 'input/data.csv';
const content = readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');
const header = lines[0].split(',');

const result: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj: any = {};
  obj.Name = values[0];
  obj.Age = parseInt(values[1], 10);
  obj.Email = values[2];
  obj.Score = parseFloat(values[3]);
  result.push(obj);
}

console.log(JSON.stringify(result));