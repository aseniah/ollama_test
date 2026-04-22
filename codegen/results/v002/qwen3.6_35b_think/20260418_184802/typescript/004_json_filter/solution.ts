import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.json');
const data = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

const result = data
  .filter((item: any) => item.active === true && item.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result, null, 2));