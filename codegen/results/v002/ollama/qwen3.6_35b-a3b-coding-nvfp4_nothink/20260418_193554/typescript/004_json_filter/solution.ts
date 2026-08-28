import * as fs from 'fs';
import * as path from 'path';

const dataPath = path.join(process.cwd(), 'input', 'data.json');
const rawData = fs.readFileSync(dataPath, 'utf-8');
const records: Array<{name: string; age: number; active: boolean; score: number}> = JSON.parse(rawData);

const filtered = records
  .filter(r => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));