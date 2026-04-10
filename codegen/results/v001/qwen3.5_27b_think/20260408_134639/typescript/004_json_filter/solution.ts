import * as fs from 'fs';

const dataPath = 'input/data.json';
const rawData = fs.readFileSync(dataPath, 'utf-8');
const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(rawData);

const filtered = data.filter(record => record.active === true && record.age >= 30);
filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));