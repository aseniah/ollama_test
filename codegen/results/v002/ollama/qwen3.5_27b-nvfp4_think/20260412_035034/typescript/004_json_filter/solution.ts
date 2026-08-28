import * as fs from 'fs';

// Read and parse the JSON file
const dataPath = 'input/data.json';
const rawData = fs.readFileSync(dataPath, 'utf8');
const records: any[] = JSON.parse(rawData);

// Filter: active is true AND age is 30 or older
const filtered = records.filter(record => 
  record.active === true && record.age >= 30
);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output as JSON array
console.log(JSON.stringify(filtered, null, 2));