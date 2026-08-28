import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join('input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf-8');
const data: Person[] = JSON.parse(rawData);

const filtered = data
  .filter(person => person.active && person.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));