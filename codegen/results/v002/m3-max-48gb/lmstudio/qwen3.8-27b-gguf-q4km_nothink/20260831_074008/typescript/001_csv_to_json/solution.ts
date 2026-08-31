import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.join('input', 'data.csv');
const content = fs.readFileSync(inputFile, 'utf-8');

const lines = content.split('\n').filter(line => line.trim() !== '');

if (lines.length < 2) {
  console.log('[]');
  process.exit(0);
}

const header = lines[0].split(',').map(h => h.trim());

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const row = lines[i].split(',').map(v => v.trim());
  if (row.length < header.length) continue;

  const name = row[0];
  const age = parseInt(row[1], 10);
  const email = row[2];
  const score = parseFloat(row[3]);

  result.push({
    Name: name,
    Age: age,
    Email: email,
    Score: score
  });
}

console.log(JSON.stringify(result, null, 2));