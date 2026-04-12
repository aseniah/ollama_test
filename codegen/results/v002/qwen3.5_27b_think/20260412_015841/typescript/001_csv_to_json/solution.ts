import * as fs from 'fs';

const filePath = 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header line
const dataLines = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = dataLines.map(line => {
  const [Name, Age, Email, Score] = line.split(',');
  return {
    Name,
    Age: parseInt(Age, 10),
    Email,
    Score: parseFloat(Score)
  };
});

console.log(JSON.stringify(result));