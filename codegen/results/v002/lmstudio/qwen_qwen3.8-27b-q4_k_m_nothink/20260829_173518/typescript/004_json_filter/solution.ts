import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const filePath = path.join('input', 'data.json');
const fileContent = fs.readFileSync(filePath, 'utf-8');
const data = JSON.parse(fileContent);

// Filter records where active is true and age >= 30
const filtered = data.filter((record: any) => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a: any, b: any) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(filtered));