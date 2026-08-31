import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.json');

// Read the file
const data = fs.readFileSync(filePath, 'utf-8');

// Parse the JSON
const records = JSON.parse(data);

// Filter records where active is true and age is 30 or older
const filteredRecords = records.filter(
  (record: { active: boolean; age: number }) => 
    record.active === true && record.age >= 30
);

// Sort by name ascending
const sortedRecords = filteredRecords.sort(
  (a: { name: string }, b: { name: string }) => 
    a.name.localeCompare(b.name)
);

// Output the result
console.log(JSON.stringify(sortedRecords));