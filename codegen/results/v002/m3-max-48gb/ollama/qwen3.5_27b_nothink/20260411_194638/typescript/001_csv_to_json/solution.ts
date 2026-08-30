import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(__dirname, '..', 'input', 'data.csv');

const fileContent = fs.readFileSync(filePath, 'utf-8');
const lines = fileContent.split('\n').filter(line => line.trim() !== '');

// Skip header
const header = lines[0].split(',');
const dataRows = lines.slice(1);

interface Person {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

const result: Person[] = [];

for (const line of dataRows) {
  const values = line.split(',');
  if (values.length !== 4) continue;

  const name = values[0].trim();
  const age = parseInt(values[1].trim(), 10);
  const email = values[2].trim();
  const score = parseFloat(values[3].trim());

  result.push({
    Name: name,
    Age: age,
    Email: email,
    Score: score
  });
}

console.log(JSON.stringify(result));