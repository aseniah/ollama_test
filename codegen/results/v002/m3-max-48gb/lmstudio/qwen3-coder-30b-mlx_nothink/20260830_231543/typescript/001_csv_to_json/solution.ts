import fs from 'fs';
import path from 'path';

// Read the CSV file
const data = fs.readFileSync(path.join('input', 'data.csv'), 'utf-8');

// Split into lines
const lines = data.trim().split('\n');

// Get header
const headers = lines[0].split(',');

// Process each row
const result = [];
for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj: any = {};
  
  for (let j = 0; j < headers.length; j++) {
    const key = headers[j].trim();
    let value = values[j].trim();
    
    // Convert types
    if (key === 'Age') {
      value = parseInt(value);
    } else if (key === 'Score') {
      value = parseFloat(value);
    }
    
    obj[key] = value;
  }
  
  result.push(obj);
}

// Output as JSON
console.log(JSON.stringify(result));