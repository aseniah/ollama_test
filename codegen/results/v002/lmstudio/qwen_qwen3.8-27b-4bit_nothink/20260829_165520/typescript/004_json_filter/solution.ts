import { readFileSync } from 'fs';

const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(readFileSync('input/data.json', 'utf-8'));

const filtered = data
  .filter((record) => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));