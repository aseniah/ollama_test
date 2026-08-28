import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputData = fs.readFileSync(path.join('input', 'data.json'), 'utf-8');
const records: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(inputData);

// Filter: active is true AND age >= 30
const filtered = records.filter(record => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output as JSON array
console.log(JSON.stringify(filtered, null, 2));