import { readFileSync } from 'fs';

const filePath = 'input/data.json';
const records = JSON.parse(readFileSync(filePath, 'utf-8')) as Array<{ name: string; age: number; active: boolean; score: number }>;

const result = records
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0));

console.log(JSON.stringify(result));