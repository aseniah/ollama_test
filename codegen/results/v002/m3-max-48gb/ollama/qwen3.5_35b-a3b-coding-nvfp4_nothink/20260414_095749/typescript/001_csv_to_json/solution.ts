import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(inputPath, 'utf-8');

const lines = content.trim().split('\n');
const headerLine = lines[0];
const dataLines = lines.slice(1);

const result: any[] = [];

for (const line of dataLines) {
  const parts = line.split(',');
  if (parts.length < 4) {
    continue;
  }
  
  const name = parts[0].trim();
  const age = parseInt(parts[1].trim(), 10);
  const email = parts[2].trim();
  const score = parseFloat(parts[3].trim());

  result.push({
    Name: name,
    Age: age,
    Email: email,
    Score: score
  });
}

console.log(JSON.stringify(result));