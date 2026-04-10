import { readFileSync, writeFileSync } from 'fs';

const rawData = JSON.parse(readFileSync('input/data.json', 'utf8'));

const filtered = rawData
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

writeFileSync('output/result.json', JSON.stringify(filtered));