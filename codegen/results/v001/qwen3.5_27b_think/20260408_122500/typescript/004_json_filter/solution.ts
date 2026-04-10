import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.json');
const data = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

const result = data
  .filter((record: { active: boolean; age: number }) => record.active && record.age >= 30)
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));