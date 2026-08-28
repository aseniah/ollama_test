import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');
const rows = lines.slice(1);

const result = rows.map((line) => {
  const fields = line.split(',');
  return {
    Name: fields[0],
    Age: parseInt(fields[1], 10),
    Email: fields[2],
    Score: parseFloat(fields[3]),
  };
});

process.stdout.write(JSON.stringify(result, null, 2) + '\n');