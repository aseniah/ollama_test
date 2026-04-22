import { readFileSync } from 'fs';

const data = JSON.parse(readFileSync('input/data.json', 'utf-8'));
const result = data
  .filter((item: { active: boolean; age: number }) => item.active === true && item.age >= 30)
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));