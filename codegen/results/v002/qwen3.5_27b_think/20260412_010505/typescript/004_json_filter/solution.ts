import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join('input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf8');
const data: Record[] = JSON.parse(rawData);

const filtered = data.filter(
  (record) => record.active === true && record.age >= 30
);

const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(sorted));