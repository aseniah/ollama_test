const fs = require('fs');

// Read the input file
const inputData = JSON.parse(fs.readFileSync('input/data.json', 'utf8'));

// Filter active records with age >= 30
const filteredData = inputData.filter(record => {
  return record.active === true && record.age >= 30;
});

// Sort by name ascending
filteredData.sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(filteredData));