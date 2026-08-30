import * as fs from 'fs';
import * as path from 'path';

// Read the file
const filePath = 'input/data.json';
const fileContent = fs.readFileSync(filePath, 'utf-8');
const data = JSON.parse(fileContent);

// Filter records where active is true AND age >= 30
const filtered = data.filter(record => record.active === true && record.age >= 30);

// Sort by name ascending
const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(sorted));