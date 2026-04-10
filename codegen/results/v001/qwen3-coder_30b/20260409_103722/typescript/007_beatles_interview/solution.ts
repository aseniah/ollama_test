import fs from 'fs';
import path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthDate: string): number {
  const birth = new Date(birthDate);
  const referenceDate = new Date('2025-07-01');
  
  let age = referenceDate.getFullYear() - birth.getFullYear();
  const monthDiff = referenceDate.getMonth() - birth.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

// Read the CSV file
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');

// Parse CSV lines
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',').map(header => header.trim());

// Process each data line
const result = lines.slice(1).map(line => {
  const values = line.split(',').map(value => value.trim());
  const person: any = {};
  
  headers.forEach((header, index) => {
    const value = values[index];
    
    // Handle different field types
    if (header === 'birth_date') {
      person[header] = value;
      person.age = calculateAge(value);
    } else if (header === 'salary') {
      person[header] = parseFloat(value);
    } else {
      person[header] = value;
    }
  });
  
  return person;
});

// Output JSON array
console.log(JSON.stringify(result, null, 2));