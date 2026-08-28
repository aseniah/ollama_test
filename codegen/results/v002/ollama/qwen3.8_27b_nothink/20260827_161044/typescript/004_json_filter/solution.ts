import * as fs from 'fs';

const data = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const filtered = data
  .filter((r: any) => r.active === true && r.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));