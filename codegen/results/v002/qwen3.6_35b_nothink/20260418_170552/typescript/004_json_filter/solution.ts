import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const filePath = path.join(__dirname, 'input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf-8');
const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(rawData);

// Filter: active is true AND age >= 30
const filtered = data.filter(record => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output as JSON array
console.log(JSON.stringify(filtered));