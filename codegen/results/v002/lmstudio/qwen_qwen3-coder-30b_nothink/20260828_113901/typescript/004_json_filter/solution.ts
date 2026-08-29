import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.json');

const rawData = fs.readFileSync(filePath, 'utf-8');
const data = JSON.parse(rawData);

const filteredData = data
  .filter((person: any) => person.active && person.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filteredData));