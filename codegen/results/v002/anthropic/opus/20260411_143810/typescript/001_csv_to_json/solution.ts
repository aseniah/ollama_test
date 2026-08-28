import fs from 'fs';
import path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.split('\n').filter(line => line.trim() !== '');
const [, ...rows] = lines;

const result = rows.map(row => {
  const [Name, Age, Email, Score] = row.split(',').map(f => f.trim());
  return {
    Name,
    Age: parseInt(Age, 10),
    Email,
    Score: parseFloat(Score),
  };
});

process.stdout.write(JSON.stringify(result));
