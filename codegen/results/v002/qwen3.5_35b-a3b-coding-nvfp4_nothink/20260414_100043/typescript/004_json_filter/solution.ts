import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'data.json');
const rawData = fs.readFileSync(inputPath, 'utf-8');
const records = JSON.parse(rawData) as Array<{ name: string; age: number; active: boolean; score: number }>;

const filtered = records.filter(record => record.active === true && record.age >= 30);
const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(sorted));