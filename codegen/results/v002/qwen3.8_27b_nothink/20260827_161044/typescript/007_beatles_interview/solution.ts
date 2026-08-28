import * as fs from 'fs';

// Read the input CSV file
const csvContent = fs.readFileSync('input/input.csv', 'utf8');
const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');
const rows = lines.slice(1).map(line => {
  const values = line.split(',');
  const obj: Record<string, string> = {};
  header.forEach((h, i) => {
    obj[h.trim()] = (values[i] || '').trim();
  });
  return obj;
});

// Helper function to parse a date string into a Date object
function parseDate(dateStr: string): Date {
  const parts = dateStr.split('/').map(Number);
  const month = parts[0] - 1;
  const day = parts[1];
  const year = parts[2];
  return new Date(year, month, day);
}

// Calculate age as of July 1, 2025
function calculateAge(birthday: string, died: string | null): number {
  const refDate = new Date(2025, 6, 1); // July 1, 2025
  const birthDate = parseDate(birthday);
  
  // Check if the person died before the reference date
  if (died && died !== 'null') {
    const deathDate = parseDate(died);
    if (deathDate < refDate) {
      // Use the death date as the reference for age calculation
      let age = deathDate.getFullYear() - birthDate.getFullYear();
      const monthDiff = deathDate.getMonth() - birthDate.getMonth();
      if (monthDiff < 0 || (monthDiff === 0 && deathDate.getDate() < birthDate.getDate())) {
        age--;
      }
      return age;
    }
  }
  
  // Still alive as of reference date, calculate age normally
  let age = refDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = refDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && refDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

// Extract first and last name from a name string
function splitName(name: string): { first: string; last: string } {
  const parts = name.split(' ');
  const first = parts[0];
  const last = parts.slice(1).join(' ');
  return { first, last };
}

// Process each row
const result = rows.map(row => {
  const { first: firstName, last: lastName } = splitName(row['Name']);
  const birthday = row['Birthday'];
  const died = row['Died'] === 'null' ? null : row['Died'];
  const age = calculateAge(birthday, died);
  
  // Convert birthday format from M/D/YYYY to YYYY-MM-DD
  const bParts = birthday.split('/').map(Number);
  const year = bParts[2].toString().padStart(4, '0');
  const month = bParts[0].toString().padStart(2, '0');
  const day = bParts[1].toString().padStart(2, '0');
  const formattedBirthday = `${year}-${month}-${day}`;
  
  // Build relatives array
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  
  const fatherName = row['Father'];
  if (fatherName && fatherName !== 'null') {
    const { first: fFirst, last: fLast } = splitName(fatherName);
    relatives.push({ FirstName: fFirst, LastName: fLast, Relationship: 'Father' });
  }
  
  const motherName = row['Mother'];
  if (motherName && motherName !== 'null') {
    const { first: mFirst, last: mLast } = splitName(motherName);
    relatives.push({ FirstName: mFirst, LastName: mLast, Relationship: 'Mother' });
  }
  
  const brotherName = row['Brother'];
  if (brotherName && brotherName !== 'null') {
    const { first: bFirst, last: bLast } = splitName(brotherName);
    relatives.push({ FirstName: bFirst, LastName: bLast, Relationship: 'Brother' });
  }
  
  const sisterName = row['Sister'];
  if (sisterName && sisterName !== 'null') {
    const { first: sFirst, last: sLast } = splitName(sisterName);
    relatives.push({ FirstName: sFirst, LastName: sLast, Relationship: 'Sister' });
  }
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: formattedBirthday,
    Age: age,
    Relatives: relatives
  };
});

// Output the JSON array to stdout
console.log(JSON.stringify(result, null, 2));