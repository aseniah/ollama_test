import * as fs from 'fs';

const data = fs.readFileSync('input/data.csv', 'utf-8');
const lines = data.split('\n').filter(line => line.trim() !== '');

// Skip header
const rows = lines.slice(1);

const result = rows.map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name.trim(),
    Age: parseInt(age.trim(), 10),
    Email: email.trim(),
    Score: parseFloat(score.trim()),
  };
});

console.log(JSON.stringify(result, null, 2));