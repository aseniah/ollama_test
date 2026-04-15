import * as fs from 'fs';

// Read the file
const content = fs.readFileSync('input/data.csv', 'utf-8');

// Parse CSV - split into lines, filter out empty lines
const lines = content.split('\n').filter(line => line.trim());

// Skip header (first line) and parse data rows
const dataRows = lines.slice(1);

const result = dataRows.map(line => {
  const values = line.split(',');
  return {
    Name: values[0],
    Age: parseInt(values[1], 10),
    Email: values[2],
    Score: parseFloat(values[3])
  };
});

// Output JSON to stdout
console.log(JSON.stringify(result));