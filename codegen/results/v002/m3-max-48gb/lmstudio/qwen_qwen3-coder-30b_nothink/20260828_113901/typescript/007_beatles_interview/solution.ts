import fs from 'fs';
import path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  const birthDate = new Date(birthday);
  const targetDate = new Date('2025-07-01');
  
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = targetDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && targetDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Helper function to parse date from MM/DD/YYYY format
function parseDate(dateStr: string): string {
  const [month, day, year] = dateStr.split('/');
  return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
}

// Read and parse CSV file
const csvData = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');

// Split into lines and parse headers
const lines = csvData.split('\n').filter(line => line.trim() !== '');
const headers = lines[0].split(',').map(h => h.trim());

// Process each data line
const result = lines.slice(1).map(line => {
  const values = line.split(',').map(v => v.trim());
  
  // Create person object
  const person: any = {
    FirstName: values[0].split(' ').filter(n => n.length > 0)[0],
    LastName: values[0].split(' ').filter(n => n.length > 0).slice(1).join(' '),
    Birthday: parseDate(values[1]),
    Age: calculateAge(parseDate(values[1])),
    Relatives: []
  };
  
  // Add relatives
  const relatives = ['Father', 'Mother', 'Brother', 'Sister'];
  for (let i = 3; i < Math.min(values.length, 7); i++) {
    if (values[i] !== 'null' && values[i].trim() !== '') {
      const relativeName = values[i].trim();
      const relationship = relatives[i - 3];
      
      // Split name into first and last name
      const parts = relativeName.split(' ');
      const relative: any = {
        FirstName: parts[0],
        LastName: parts.slice(1).join(' '),
        Relationship: relationship
      };
      
      person.Relatives.push(relative);
    }
  }
  
  return person;
});

// Output as JSON
console.log(JSON.stringify(result, null, 2));