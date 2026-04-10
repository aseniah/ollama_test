import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf-8');
const data: any[] = JSON.parse(rawData);

const result = data
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));