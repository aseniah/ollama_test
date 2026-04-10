import * as fs from 'fs';

const data = JSON.parse(fs.readFileSync('input/data.json', 'utf8'));

// Filter records: active === true AND age >= 30
const filteredData = data.filter((record) => record.active && record.age >= 30);

// Sort by name ascending
filteredData.sort((a, b) => a.name.localeCompare(b.name));

// Output as JSON to stdout
console.log(JSON.stringify(filteredData, null, 2));