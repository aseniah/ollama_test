import * as fs from 'fs';

// Read the input file
const data = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

// Filter: active is true AND age >= 30
const filtered = data.filter((record: any) => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a: any, b: any) => a.name.localeCompare(b.name));

// Output the result as JSON
console.log(JSON.stringify(filtered));