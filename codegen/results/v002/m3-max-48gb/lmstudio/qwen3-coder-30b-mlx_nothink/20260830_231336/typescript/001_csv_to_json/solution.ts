import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = join('input', 'data.csv');
const fileContent = readFileSync(filePath, 'utf-8');

const lines = fileContent.trim().split('\n');
const headers = lines[0].split(',');

const result = lines.slice(1).map(line => {
  const values = line.split(',');
  const obj: any = {};
  
  headers.forEach((header, index) => {
    const key = header.trim();
    const value = values[index]?.trim();
    
    switch (key) {
      case 'Name':
        obj[key] = value;
        break;
      case 'Age':
        obj[key] = parseInt(value, 10);
        break;
      case 'Email':
        obj[key] = value;
        break;
      case 'Score':
        obj[key] = parseFloat(value);
        break;
    }
  });
  
  return obj;
});

console.log(JSON.stringify(result));