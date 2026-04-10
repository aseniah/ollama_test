import * as fs from 'fs';

const rawData = fs.readFileSync('input/data.json', 'utf8');
const data: any[] = JSON.parse(rawData);

const filtered = data.filter(record => 
  record.active === true && 
  (record.age >= 30)
);

filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));