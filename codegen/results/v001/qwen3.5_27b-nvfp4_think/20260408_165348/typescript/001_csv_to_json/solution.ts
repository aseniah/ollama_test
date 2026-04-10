import * as fs from 'fs';
import * as path from 'path';

const filePath = process.argv[2] || 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header (first line)
const dataLines = lines.slice(1);

interface DataRow {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

const result: DataRow[] = [];

for (const line of dataLines) {
  const [name, ageStr, email, scoreStr] = line.split(',');
  
  if (!name || !ageStr || !email || !scoreStr) continue;

  result.push({
    Name: name.trim(),
    Age: parseInt(ageStr.trim(), 10),
    Email: email.trim(),
    Score: parseFloat(scoreStr.trim())
  });
}

console.log(JSON.stringify(result));