import * as fs from 'fs';

// Read the CSV file
const filePath = 'input/data.csv';
const csvContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines and filter out empty lines
const lines = csvContent.split('\n').filter(line => line.trim() !== '');

// Skip the header row
const dataLines = lines.slice(1);

// Parse each line into an object
const result: Array<{
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}> = dataLines.map(line => {
  // Split by comma (Name field contains spaces but not commas in this data)
  const parts = line.split(',');

  return {
    Name: parts[0],
    Age: parseInt(parts[1], 10),
    Email: parts[2],
    Score: parseFloat(parts[3]),
  };
});

// Output as JSON to stdout
console.log(JSON.stringify(result, null, 2));