import * as fs from 'fs';

const filePath = 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split(/\r?\n/);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const [name, ageStr, email, scoreStr] = lines[i].split(',');
  result.push({
    Name: name,
    Age: parseInt(ageStr, 10),
    Email: email,
    Score: parseFloat(scoreStr)
  });
}

console.log(JSON.stringify(result));