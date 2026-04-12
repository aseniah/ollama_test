import * as fs from 'fs';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const data: Person[] = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const result = data
  .filter(p => p.active && p.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));