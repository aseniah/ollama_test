import * as fs from 'fs';
import * as path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthDate: string): number {
  const birth = new Date(birthDate);
  const reference = new Date('2025-07-01');
  
  let age = reference.getFullYear() - birth.getFullYear();
  const monthDiff = reference.getMonth() - birth.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && reference.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

// Read the CSV file
const csvPath = path.join('input', 'input.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

// Parse CSV content
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',').map(header => header.trim());

const result = lines.slice(1).map(line => {
  const values = line.split(',').map(value => value.trim());
  const person: any = {};
  
  headers.forEach((header, index) => {
    const value = values[index];
    if (header === 'birth_date') {
      person.age = calculateAge(value);
    }
    person[header] = value;
  });
  
  return person;
});

console.log(JSON.stringify(result, null, 2));