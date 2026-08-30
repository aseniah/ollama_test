import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.split('\n').filter(line => line.trim() !== '');
const header = lines[0].split(',');
const dataRows = lines.slice(1);

const result = dataRows.map(row => {
  const cols = row.split(',');
  return {
    Name: cols[0].trim(),
    Age: parseInt(cols[1].trim(), 10),
    Email: cols[2].trim(),
    Score: parseFloat(cols[3].trim())
  };
});

console.log(JSON.stringify(result, null, 2));