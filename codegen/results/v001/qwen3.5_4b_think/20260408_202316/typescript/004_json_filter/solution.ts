import * as fs from 'fs';

// Read the JSON file
const data = JSON.parse(fs.readFileSync('input/data.json', 'utf8'));

// Filter and sort
const result = data
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(result));