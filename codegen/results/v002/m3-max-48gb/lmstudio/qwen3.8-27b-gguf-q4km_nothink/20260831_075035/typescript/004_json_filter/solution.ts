import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'data.json');
const raw = fs.readFileSync(inputPath, 'utf-8');
const data: Array<{name: string; age: number; active: boolean; score: number}> = JSON.parse(raw);

const filtered = data
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));