import fs from 'fs';

const data = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));
const result = data
  .filter((r: any) => r.active === true && r.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));
process.stdout.write(JSON.stringify(result));
