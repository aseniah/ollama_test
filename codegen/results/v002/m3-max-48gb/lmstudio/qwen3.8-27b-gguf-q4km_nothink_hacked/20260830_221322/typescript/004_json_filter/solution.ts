import * as fs from 'fs';

const dataPath = 'input/data.json';

// Read and parse the JSON file
const raw = fs.readFileSync(dataPath, 'utf-8');
const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(raw);

// Filter: active is true AND age >= 30
const filtered = data.filter((record) => record.active === true && record.age >= 30);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output the JSON array to stdout
console.log(JSON.stringify(filtered));