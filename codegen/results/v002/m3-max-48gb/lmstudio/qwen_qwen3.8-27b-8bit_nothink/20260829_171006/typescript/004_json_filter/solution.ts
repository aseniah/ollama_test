import * as fs from 'fs';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const data: Record[] = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const filtered = data
  .filter(r => r.active && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));