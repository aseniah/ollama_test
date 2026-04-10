import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');

// Parse CSV content
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');

const result = lines.slice(1).map(line => {
  const values = line.split(',');
  const row: any = {};
  
  headers.forEach((header, index) => {
    const value = values[index]?.trim();
    row[header] = value;
  });
  
  // Calculate age as of July 1, 2025
  if (row['birth_date']) {
    const birthDate = new Date(row['birth_date']);
    const endDate = new Date('2025-07-01');
    
    let age = endDate.getFullYear() - birthDate.getFullYear();
    const monthDiff = endDate.getMonth() - birthDate.getMonth();
    
    if (monthDiff < 0 || (monthDiff === 0 && endDate.getDate() < birthDate.getDate())) {
      age--;
    }
    
    row['age'] = age;
  }
  
  return row;
});

console.log(JSON.stringify(result, null, 2));