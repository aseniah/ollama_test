import fs from 'fs';
import path from 'path';

// Read the input file
const inputPath = path.join(process.cwd(), 'input', 'data.json');
const data = JSON.parse(fs.readFileSync(inputPath, 'utf-8'));

// Filter records: active is true AND age >= 30
const filtered = data.filter(record => 
  record.active === true && record.age >= 30
);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout as JSON array
console.log(JSON.stringify(filtered));