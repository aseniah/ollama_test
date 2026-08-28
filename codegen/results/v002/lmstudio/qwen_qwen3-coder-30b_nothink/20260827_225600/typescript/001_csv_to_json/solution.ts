import fs from 'fs';
import path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const fileContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines and process
const lines = fileContent.trim().split('\n');

// Skip header line and process each row
const result = lines.slice(1).map(line => {
  const [name, age, email, score] = line.split(',');
  return {
    Name: name.trim(),
    Age: parseInt(age.trim(), 10),
    Email: email.trim(),
    Score: parseFloat(score.trim())
  };
});

// Output as JSON
console.log(JSON.stringify(result));