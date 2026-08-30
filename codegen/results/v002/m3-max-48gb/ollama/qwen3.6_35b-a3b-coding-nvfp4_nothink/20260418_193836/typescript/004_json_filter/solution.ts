import * as fs from 'fs';
import * as path from 'path';

const data = JSON.parse(fs.readFileSync(path.join(__dirname, 'input/data.json'), 'utf-8'));

const result = data
  .filter((record: { name: string; age: number; active: boolean; score: number }) => record.active === true && record.age >= 30)
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));