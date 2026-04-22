import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

const result = lines.slice(1).map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name,
    Age: parseInt(age, 10),
    Email: email,
    Score: parseFloat(score)
  };
});

console.log(JSON.stringify(result));