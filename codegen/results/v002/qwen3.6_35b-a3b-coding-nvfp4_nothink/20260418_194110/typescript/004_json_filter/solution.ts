import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.json');
const data = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

const filtered = data
  .filter((record: { name: string; age: number; active: boolean; score: number }) => record.active === true && record.age >= 30)
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));