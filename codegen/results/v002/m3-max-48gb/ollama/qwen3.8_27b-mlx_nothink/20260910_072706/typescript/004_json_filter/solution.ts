import { readFileSync, writeFileSync } from 'fs';
import { join } from 'path';

// Read the input file
const inputPath = join('input', 'data.json');
const data = JSON.parse(readFileSync(inputPath, 'utf-8'));

// Filter: active is true and age >= 30
const filtered = data.filter((record: any) => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a: any, b: any) => a.name.localeCompare(b.name));

// Output the JSON array to stdout
process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');