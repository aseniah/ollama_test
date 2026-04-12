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

// Helper function to parse date from MM/DD/YYYY format
function parseDate(dateStr: string): string {
  const [month, day, year] = dateStr.split('/');
  return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
}

// Read and parse CSV file
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const lines = csvContent.trim().split('\n');

// Skip header line
const dataLines = lines.slice(1);

const result = dataLines.map(line => {
  const [name, birthday, died, father, mother, brother, sister] = line.split(',');
  
  // Parse name into FirstName and LastName
  const nameParts = name.split(' ');
  const firstName = nameParts[0];
  const lastName = nameParts.slice(1).join(' ');
  
  // Parse birthday
  const parsedBirthday = parseDate(birthday);
  
  // Calculate age
  const age = calculateAge(parsedBirthday);
  
  // Build relatives array
  const relatives = [];
  
  if (father !== 'null') {
    const [fFirstName, fLastName] = father.split(' ');
    relatives.push({
      FirstName: fFirstName,
      LastName: fLastName,
      Relationship: 'Father'
    });
  }
  
  if (mother !== 'null') {
    const [mFirstName, mLastName] = mother.split(' ');
    relatives.push({
      FirstName: mFirstName,
      LastName: mLastName,
      Relationship: 'Mother'
    });
  }
  
  if (brother !== 'null') {
    const [bFirstName, bLastName] = brother.split(' ');
    relatives.push({
      FirstName: bFirstName,
      LastName: bLastName,
      Relationship: 'Brother'
    });
  }
  
  if (sister !== 'null') {
    const [sFirstName, sLastName] = sister.split(' ');
    relatives.push({
      FirstName: sFirstName,
      LastName: sLastName,
      Relationship: 'Sister'
    });
  }
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: parsedBirthday,
    Age: age,
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));