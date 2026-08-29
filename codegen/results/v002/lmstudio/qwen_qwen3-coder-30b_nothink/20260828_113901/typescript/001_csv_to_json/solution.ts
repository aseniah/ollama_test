import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.csv');

const data = fs.readFileSync(filePath, 'utf-8');

const lines = data.trim().split('\n');

// Skip header
const rows = lines.slice(1);

const result = rows.map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name,
    Age: parseInt(age, 10),
    Email: email,
    Score: parseFloat(score)
  };
});

console.log(JSON.stringify(result));