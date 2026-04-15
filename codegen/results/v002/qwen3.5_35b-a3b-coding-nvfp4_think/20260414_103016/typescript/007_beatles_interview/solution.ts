import * as fs from 'fs';

// Read the CSV file
const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

// Skip header line
const dataLines = lines.slice(1);

// Parse date from M/D/YYYY to YYYY-MM-DD format
function parseDate(dateStr: string): string {
  const parts = dateStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
}

// Calculate age as of July 1, 2025
function calculateAge(birthdayStr: string): number {
  const birthday = new Date(birthdayStr);
  const referenceDate = new Date('2025-07-01');
  
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  
  // Adjust if birthday hasn't occurred yet in the reference year
  const birthdayMonth = birthday.getMonth();
  const birthdayDay = birthday.getDate();
  const referenceMonth = referenceDate.getMonth();
  const referenceDay = referenceDate.getDate();
  
  if (birthdayMonth > referenceMonth || 
      (birthdayMonth === referenceMonth && birthdayDay > referenceDay)) {
    age--;
  }
  
  return age;
}

// Parse name to extract first and last name
function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(' ');
  return {
    firstName: parts[0],
    lastName: parts[parts.length - 1]
  };
}

// Parse relative name to extract first and last name
function parseRelativeName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(' ');
  return {
    firstName: parts[0],
    lastName: parts[parts.length - 1]
  };
}

// Build the output array
const output: any[] = [];

for (const line of dataLines) {
  const columns = line.split(',');
  
  const { firstName, lastName } = parseName(columns[0]);
  const birthdayStr = parseDate(columns[1]);
  const age = calculateAge(birthdayStr);
  
  const relatives: any[] = [];
  
  // Father
  if (columns[3] && columns[3].trim() !== 'null') {
    const { firstName: fatherFirst, lastName: fatherLast } = parseRelativeName(columns[3]);
    relatives.push({
      FirstName: fatherFirst,
      LastName: fatherLast,
      Relationship: "Father"
    });
  }
  
  // Mother
  if (columns[4] && columns[4].trim() !== 'null') {
    const { firstName: motherFirst, lastName: motherLast } = parseRelativeName(columns[4]);
    relatives.push({
      FirstName: motherFirst,
      LastName: motherLast,
      Relationship: "Mother"
    });
  }
  
  // Brother
  if (columns[5] && columns[5].trim() !== 'null') {
    const { firstName: brotherFirst, lastName: brotherLast } = parseRelativeName(columns[5]);
    relatives.push({
      FirstName: brotherFirst,
      LastName: brotherLast,
      Relationship: "Brother"
    });
  }
  
  // Sister
  if (columns[6] && columns[6].trim() !== 'null') {
    const { firstName: sisterFirst, lastName: sisterLast } = parseRelativeName(columns[6]);
    relatives.push({
      FirstName: sisterFirst,
      LastName: sisterLast,
      Relationship: "Sister"
    });
  }
  
  output.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives
  });
}

// Output JSON to stdout
console.log(JSON.stringify(output, null, 2));