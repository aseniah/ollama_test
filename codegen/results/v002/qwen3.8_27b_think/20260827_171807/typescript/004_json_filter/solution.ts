import * as fs from 'fs';

const data: { name: string; age: number; active: boolean; score: number }[] =
  JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const result = data
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(result, null, 2) + '\n');