import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const fileContent = fs.readFileSync(filePath, 'utf-8');

// Split into lines
const lines = fileContent.trim().split('\n');

// Get header row
const headers = lines[0].split(',');

// Process each data row
const result = lines.slice(1).map(line => {
  const values = line.split(',');
  const obj: any = {};
  
  headers.forEach((header, index) => {
    const value = values[index].trim();
    const key = header.trim();
    
    // Convert types based on header name
    switch (key) {
      case 'Age':
        obj[key] = parseInt(value, 10);
        break;
      case 'Score':
        obj[key] = parseFloat(value);
        break;
      default:
        obj[key] = value;
    }
  });
  
  return obj;
});

// Output as JSON
console.log(JSON.stringify(result));