import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, 'data.csv');

fs.readFile(inputFilePath, 'utf-8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  const lines = data.split('\n');
  const header = lines[0].split(',');
  const rows = lines.slice(1).map(row => row.split(','));

  const result: any[] = rows.map(([Name, Age, Email, Score]) => ({
    Name,
    Age: parseInt(Age),
    Email,
    Score: parseFloat(Score)
  }));

  console.log(JSON.stringify(result));
});