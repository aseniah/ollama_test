import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(process.cwd(), 'input', 'data.csv');
const output = [];

const csvContent = fs.readFileSync(inputFilePath, 'utf8');
const lines = csvContent.trim().split('\n');

const headers = lines[0].split(',');
const data = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj = {};

  headers.forEach((header, index) => {
    const value = values[index];
    if (header === 'Name') {
      obj.Name = value;
    } else if (header === 'Age') {
      obj.Age = parseInt(value, 10);
    } else if (header === 'Email') {
      obj.Email = value;
    } else if (header === 'Score') {
      obj.Score = parseFloat(value);
    }
  });

  data.push(obj);
}

console.log(JSON.stringify(data));