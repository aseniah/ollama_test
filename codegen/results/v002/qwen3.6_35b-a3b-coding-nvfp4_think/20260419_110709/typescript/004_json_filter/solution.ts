import { readFileSync } from 'fs';

const data = JSON.parse(readFileSync('input/data.json', 'utf-8'));
const result = data
  .filter((item: any) => item.active === true && item.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));