import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const csvPath = path.join('input', 'data.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

// Parse CSV - split by lines
const lines = csvContent.trim().split('\n');
const header = lines[0];
const dataLines = lines.slice(1);

// Parse each data line into objects
const result = dataLines.map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name,
    Age: parseInt(age),
    Email: email,
    Score: parseFloat(score)
  };
});

// Output as valid JSON array
console.log(JSON.stringify(result));