import { readFileSync } from 'fs';

const csvContent = readFileSync('input/data.csv', 'utf-8');

const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');

// Define the expected fields with their types
interface DataItem {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

const results: DataItem[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const item: DataItem = {
    Name: values[0].trim(),
    Age: parseInt(values[1].trim(), 10),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim())
  };
  results.push(item);
}

// Output as JSON to stdout
console.log(JSON.stringify(results, null, 2));