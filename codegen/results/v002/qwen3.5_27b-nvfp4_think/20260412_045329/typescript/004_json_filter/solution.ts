import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join('input', 'data.json');
const data = JSON.parse(fs.readFileSync(filePath, 'utf-8')) as Record[];

const filtered = data.filter(record => 
  record.active === true && record.age >= 30
);

filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));