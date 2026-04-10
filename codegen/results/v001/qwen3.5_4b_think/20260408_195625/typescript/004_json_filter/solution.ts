const fs = require('fs');
const path = require('path');

// Read the input file
const data = JSON.parse(fs.readFileSync(path.join(__dirname, 'input', 'data.json'), 'utf-8'));

// Filter records where active is true and age >= 30, then sort by name
const result = data
  .filter(record => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

// Output JSON array to stdout
console.log(JSON.stringify(result));