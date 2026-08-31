import fs from 'fs';
import path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const fileContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines
const lines = fileContent.trim().split('\n');

// Get header row
const headers = lines[0].split(',');

// Process each data row
const result = [];
for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj: any = {};
  
  headers.forEach((header, index) => {
    const key = header.trim();
    const value = values[index]?.trim();
    
    // Convert types
    if (key === 'Age') {
      obj[key] = parseInt(value, 10);
    } else if (key === 'Score') {
      obj[key] = parseFloat(value);
    } else {
      obj[key] = value;
    }
  });
  
  result.push(obj);
}

// Output as JSON
console.log(JSON.stringify(result));