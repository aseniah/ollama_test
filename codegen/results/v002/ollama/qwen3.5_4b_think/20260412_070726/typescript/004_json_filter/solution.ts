import * as fs from 'fs';
import * as path from 'path';

const dataPath = 'input/data.json';
const data = JSON.parse(fs.readFileSync(dataPath, 'utf-8'));

// Filter records where active is true and age is 30 or older
const filtered = data.filter(record => 
  record.active === true && record.age >= 30
);

// Sort by name ascending
const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(sorted));