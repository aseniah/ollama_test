import * as fs from 'fs';

const data = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const filtered = data
  .filter((record: { active: boolean; age: number }) => record.active === true && record.age >= 30)
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(filtered, null, 0));