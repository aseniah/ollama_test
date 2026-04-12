import * as fs from 'fs';

// Read the input file
const inputPath = './input/data.json';
const data = JSON.parse(fs.readFileSync(inputPath, 'utf-8'));

// Filter records where active is true AND age is 30 or older
const filtered = data.filter(record => record.active && record.age >= 30);

// Sort the filtered records by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(filtered));