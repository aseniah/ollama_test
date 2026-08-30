import * as fs from 'fs';
import * as path from 'path';

const filePath = 'input/data.json';
const fileContent = fs.readFileSync(filePath, 'utf-8');
const records = JSON.parse(fileContent);

const filtered = records.filter(record => 
  record.active === true && record.age >= 30
);

const sorted = filtered.sort((a, b) => 
  a.name.localeCompare(b.name)
);

console.log(JSON.stringify(sorted));