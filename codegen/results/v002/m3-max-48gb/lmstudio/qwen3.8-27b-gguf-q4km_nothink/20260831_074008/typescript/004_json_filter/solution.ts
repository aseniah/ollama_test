import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const inputPath = path.resolve('input/data.json');
const data: Person[] = JSON.parse(fs.readFileSync(inputPath, 'utf-8'));

const filtered = data
  .filter((p) => p.active === true && p.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered, null, 2));