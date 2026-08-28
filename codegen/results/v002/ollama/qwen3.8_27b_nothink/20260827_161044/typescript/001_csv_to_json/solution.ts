import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const data = fs.readFileSync(filePath, 'utf8');

// Split into lines
const lines = data.trim().split('\n');

// Skip header
const rows = lines.slice(1);

// Parse each row
const result = rows.map((row) => {
  const [Name, Age, Email, Score] = row.split(',');
  return {
    Name: Name.trim(),
    Age: parseInt(Age.trim(), 10),
    Email: Email.trim(),
    Score: parseFloat(Score.trim())
  };
});

// Output as JSON
console.log(JSON.stringify(result));