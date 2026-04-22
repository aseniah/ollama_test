import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.resolve('input/data.json');
const rawData = fs.readFileSync(inputPath, 'utf-8');
const records: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(rawData);

// Filter: active is true and age is 30 or older
const filtered = records.filter(r => r.active === true && r.age >= 30);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output as JSON array
console.log(JSON.stringify(filtered, null, 2));