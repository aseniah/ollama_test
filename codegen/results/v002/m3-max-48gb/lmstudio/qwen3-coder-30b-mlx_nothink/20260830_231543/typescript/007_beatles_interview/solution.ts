import fs from 'fs';
import path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  const birthDate = new Date(birthday);
  const referenceDate = new Date('2025-07-01');
  
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  // Adjust age if birthday hasn't occurred yet this year
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Helper function to parse date from MM/DD/YYYY format to YYYY-MM-DD
function parseDate(dateStr: string): string {
  const [month, day, year] = dateStr.split('/');
  return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
}

// Helper function to parse names from CSV format (removing any trailing nulls)
function parseNames(namesStr: string): string[] {
  if (!namesStr || namesStr === 'null') return [];
  return namesStr.split(',').map(name => name.trim()).filter(name => name && name !== 'null');
}

// Read and parse the CSV file
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const lines = csvContent.trim().split('\n');

// Skip header line
const dataLines = lines.slice(1);

const result = dataLines.map(line => {
  const [name, birthday, died, father, mother, brother, sister] = line.split(',');
  
  // Parse name into FirstName and LastName
  const [firstName, ...lastNameParts] = name.split(' ').filter(part => part.trim() !== '');
  const lastName = lastNameParts.join(' ') || '';
  
  // Parse birthday
  const birthDate = parseDate(birthday);
  
  // Calculate age as of July 1, 2025
  const age = calculateAge(birthDate);
  
  // Collect relatives
  const relatives = [];
  
  // Add father if present
  if (father && father !== 'null') {
    const [fName, fNameLast] = father.split(' ');
    relatives.push({
      FirstName: fName,
      LastName: fNameLast || '',
      Relationship: "Father"
    });
  }
  
  // Add mother if present
  if (mother && mother !== 'null') {
    const [mName, mNameLast] = mother.split(' ');
    relatives.push({
      FirstName: mName,
      LastName: mNameLast || '',
      Relationship: "Mother"
    });
  }
  
  // Add brother if present
  if (brother && brother !== 'null') {
    const [bName, bNameLast] = brother.split(' ');
    relatives.push({
      FirstName: bName,
      LastName: bNameLast || '',
      Relationship: "Brother"
    });
  }
  
  // Add sister if present
  if (sister && sister !== 'null') {
    const [sName, sNameLast] = sister.split(' ');
    relatives.push({
      FirstName: sName,
      LastName: sNameLast || '',
      Relationship: "Sister"
    });
  }
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthDate,
    Age: age,
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));