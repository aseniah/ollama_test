import fs from 'fs';

const filePath = 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');

const lines = content.trim().split('\n');
const header = lines[0]; // Skip header
const data = lines.slice(1);

const results = data.map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    name: name.trim(),
    age: parseInt(age.trim(), 10),
    email: email.trim(),
    score: parseFloat(score.trim()),
  };
});

console.log(JSON.stringify(results, null, 2));