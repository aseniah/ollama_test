import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf-8');

const data = JSON.parse(rawData);

const filtered = data
  .filter((item: any) => item.active === true && item.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));