import fs from 'fs';
import path from 'path';

// Read the input file
const data = fs.readFileSync(path.join('input', 'data.json'), 'utf-8');

// Parse the JSON
const records = JSON.parse(data);

// Filter records where active is true and age is 30 or older
const filteredRecords = records.filter((record: any) => 
  record.active === true && record.age >= 30
);

// Sort by name ascending
filteredRecords.sort((a: any, b: any) => {
  if (a.name < b.name) return -1;
  if (a.name > b.name) return 1;
  return 0;
});

// Output the result
console.log(JSON.stringify(filteredRecords));