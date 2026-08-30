import * as fs from 'fs';

interface DataRow {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

const filePath = 'input/data.csv';
const fileContent = fs.readFileSync(filePath, 'utf-8');
const lines = fileContent.trim().split('\n');

// Skip header row
const dataRows = lines.slice(1);

const result: DataRow[] = dataRows.map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name,
    Age: parseInt(age, 10),
    Email: email,
    Score: parseFloat(score)
  };
});

console.log(JSON.stringify(result));