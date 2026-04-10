import { readFileSync } from 'fs';
import * as path from 'path';

// Read the file content
const filePath = path.join(__dirname, 'input', 'data.json');
const fileContent = readFileSync(filePath, 'utf8');

// Parse the JSON array
const data: { name: string; age: number; active: boolean; score: number }[] = JSON.parse(fileContent);

// Filter and sort the data
const filteredData = data.filter(record => record.active && record.age >= 30);
const sortedData = filteredData.sort((a, b) => a.name.localeCompare(b.name));

// Output the result as JSON array
console.log(JSON.stringify(sortedData, null, 2));