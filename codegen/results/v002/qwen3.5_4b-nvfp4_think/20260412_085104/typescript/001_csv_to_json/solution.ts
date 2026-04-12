import { readFileSync } from 'fs';
import { resolve } from 'path';

const filePath = resolve(__dirname, 'input/data.csv');
const content = readFileSync(filePath, 'utf-8');

const lines = content.trim().split('\n');
const dataRows = lines.slice(1);

const result: any[] = [];

for (const row of dataRows) {
  const values = row.split(',');
  const name = values[0];
  const age = parseInt(values[1]);
  const email = values[2];
  const score = parseFloat(values[3]);

  result.push({
    Name: name,
    Age: age,
    Email: email,
    Score: score,
  });
}

console.log(JSON.stringify(result, null, 2));