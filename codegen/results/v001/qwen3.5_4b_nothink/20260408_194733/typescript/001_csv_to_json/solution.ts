import { readFileSync } from 'fs';

const content = readFileSync('input/data.csv', 'utf-8');
const lines = content.trim().split('\n');

// Skip header
const dataLines = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = dataLines.map((line) => {
  // Assuming the format is: Name,Age,Email,Score
  const [name, age, email, score] = line.split(',');
  
  return {
    Name: name.trim(),
    Age: parseInt(age.trim(), 10),
    Email: email.trim(),
    Score: parseFloat(score.trim()),
  };
});

console.log(JSON.stringify(result));