import { readFileSync } from 'fs';

const dataFile = './input/data.json';
const content = readFileSync(dataFile, 'utf-8');
const records: any[] = JSON.parse(content);

const result = records
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result, null, 2));