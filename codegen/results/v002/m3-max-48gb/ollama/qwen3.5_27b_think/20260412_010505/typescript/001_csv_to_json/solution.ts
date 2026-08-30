import * as fs from 'fs';

const filePath = 'input/data.csv';

// Read the file content
const fileContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines and filter out empty lines
const lines = fileContent.split('\n').filter(line => line.trim() !== '');

// Skip the header (first line)
const dataLines = lines.slice(1);

// Parse each data line into an object
const results = dataLines.map(line => {
  const parts = line.split(',');
  
  return {
    Name: parts[0].trim(),
    Age: parseInt(parts[1].trim(), 10),
    Email: parts[2].trim(),
    Score: parseFloat(parts[3].trim())
  };
});

// Output as JSON to stdout
console.log(JSON.stringify(results));