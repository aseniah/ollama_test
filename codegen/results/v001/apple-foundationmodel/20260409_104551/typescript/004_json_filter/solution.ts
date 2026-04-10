import fs from 'fs';
import path from 'path';
import { readFileSync } from 'fs';

// Define the input file path
const inputFilePath = path.join(__dirname, 'input', 'data.json');

// Read the JSON file
const jsonData = readFileSync(inputFilePath, 'utf8');

// Parse JSON data
const data: any[] = JSON.parse(jsonData);

// Filter records where 'active' is true and 'age' is 30 or older
const filteredData = data.filter(record => {
  const { active, age } = record;
  return active && age >= 30;
});

// Sort the filtered records by 'name' in ascending order
const sortedData = filteredData.sort((a, b) => {
  const { name } = a;
  const { name } = b;
  return name.localeCompare(name);
});

// Output the sorted records to stdout
console.log(JSON.stringify(sortedData, null, 2));