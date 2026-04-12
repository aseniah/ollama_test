import { readFileSync } from 'fs';
import { join } from 'path';

// Read the input file
const filePath = join('input', 'data.json');
const fileContent = readFileSync(filePath, 'utf-8');

// Parse the JSON
const data = JSON.parse(fileContent);

// Filter records where active is true and age is 30 or older
const filteredData = data.filter((record: any) => 
  record.active === true && record.age >= 30
);

// Sort by name ascending
filteredData.sort((a: any, b: any) => {
  if (a.name < b.name) return -1;
  if (a.name > b.name) return 1;
  return 0;
});

// Output the result
console.log(JSON.stringify(filteredData));