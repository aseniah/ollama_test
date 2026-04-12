import fs from 'fs';
import path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');
const rows = lines.slice(1).map(line => {
  const [Name, Age, Email, Score] = line.split(',');
  return {
    Name: Name.trim(),
    Age: parseInt(Age.trim(), 10),
    Email: Email.trim(),
    Score: parseFloat(Score.trim()),
  };
});
console.log(JSON.stringify(rows));
