import * as fs from 'fs';
import * as path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  const birthDate = new Date(birthday);
  const referenceDate = new Date('2025-07-01');
  
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  // Adjust if birthday hasn't occurred yet this year
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Parse CSV data
const csvData = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const lines = csvData.trim().split('\n');

// Skip header line
const headers = lines[0].split(',');
const rows = lines.slice(1);

// Process each row
const result = rows.map(row => {
  const values = row.split(',');
  const person: any = {};
  
  // Map headers to values
  headers.forEach((header, index) => {
    const value = values[index].trim();
    if (value !== 'null' && value !== '') {
      person[header] = value;
    }
  });
  
  // Transform data
  const firstName = person.Name.split(' ')[0];
  const lastName = person.Name.split(' ').slice(-1)[0];
  const birthday = person.Birthday;
  
  // Convert birthday format from MM/DD/YYYY to YYYY-MM-DD
  const [month, day, year] = birthday.split('/');
  const formattedBirthday = `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
  
  // Create relatives array
  const relatives = [];
  
  if (person.Father) {
    const [fFirstName, fLastName] = person.Father.split(' ');
    relatives.push({
      FirstName: fFirstName,
      LastName: fLastName,
      Relationship: "Father"
    });
  }
  
  if (person.Mother) {
    const [mFirstName, mLastName] = person.Mother.split(' ');
    relatives.push({
      FirstName: mFirstName,
      LastName: mLastName,
      Relationship: "Mother"
    });
  }
  
  if (person.Brother) {
    const [bFirstName, bLastName] = person.Brother.split(' ');
    relatives.push({
      FirstName: bFirstName,
      LastName: bLastName,
      Relationship: "Brother"
    });
  }
  
  if (person.Sister) {
    const [sFirstName, sLastName] = person.Sister.split(' ');
    relatives.push({
      FirstName: sFirstName,
      LastName: sLastName,
      Relationship: "Sister"
    });
  }
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: formattedBirthday,
    Age: calculateAge(formattedBirthday),
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));