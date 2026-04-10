const fs = require('fs');
const path = require('path');

const inputPath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(inputPath, 'utf-8');

const lines = content.trim().split('\n');
if (lines.length === 0) {
  process.stdout.write('[]');
  process.exit(0);
}

const data = [];
const header = lines[0].split(',').map((field) => field.trim());

const requiredFields = ['name', 'age', 'email', 'score'];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i];
  if (!line.trim()) continue;

  const values = line.split(',').map((value) => value.trim());
  
  const obj = {
    name: values[0],
    age: parseInt(values[1], 10),
    email: values[2],
    score: parseFloat(values[3]),
  };

  data.push(obj);
}

process.stdout.write(JSON.stringify(data));