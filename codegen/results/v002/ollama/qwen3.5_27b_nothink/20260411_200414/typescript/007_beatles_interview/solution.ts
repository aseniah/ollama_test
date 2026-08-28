import * as fs from 'fs';
import * as path from 'path';

// Calculate age as of July 1, 2025
function calculateAge(birthdayStr: string): number {
  const refDate = new Date('2025-07-01');
  
  // Parse the birthday string (MM/DD/YYYY)
  const parts = birthdayStr.split('/');
  const month = parseInt(parts[0], 10) - 1; // JS months are 0-indexed
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  
  const birthDate = new Date(year, month, day);
  
  let age = refDate.getFullYear() - birthDate.getFullYear();
  
  // Adjust if birthday hasn't occurred yet in the reference year
  const monthDay = refDate.getMonth() * 32 + refDate.getDate();
  const birthMonthDay = birthDate.getMonth() * 32 + birthDate.getDate();
  
  if (monthDay < birthMonthDay) {
    age--;
  }
  
  return age;
}

// Parse a full name into FirstName and LastName
function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  return {
    firstName: parts[0],
    lastName: parts.slice(1).join(' ')
  };
}

// Parse a relative name string (e.g., "Alfred Lennon" or "null")
function parseRelativeName(name: string): string[] {
  if (!name || name.toLowerCase() === 'null') {
    return [];
  }
  return name.split(/\s+/);
}

// Transform the name parts into the required format for relatives
// Expected: FirstName and LastName (first word is first, rest is last)
function formatRelativeName(parts: string[]) {
  if (parts.length === 0) return null;
  return {
    FirstName: parts[0],
    LastName: parts.slice(1).join(' ') || ''
  };
}

// Main logic
function main() {
  const inputPath = path.join('input', 'input.csv');
  
  let content: string;
  try {
    content = fs.readFileSync(inputPath, 'utf-8');
  } catch (err) {
    console.error('Error reading input file:', (err as Error).message);
    process.exit(1);
  }

  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  const headers = lines[0].split(',');
  
  const result: any[] = [];

  // Indices for columns
  const idxName = headers.indexOf('Name');
  const idxBirthday = headers.indexOf('Birthday');
  const idxDied = headers.indexOf('Died');
  const idxFather = headers.indexOf('Father');
  const idxMother = headers.indexOf('Mother');
  const idxBrother = headers.indexOf('Brother');
  const idxSister = headers.indexOf('Sister');

  for (let i = 1; i < lines.length; i++) {
    // Simple CSV parsing (assuming no commas inside quoted fields based on sample)
    const values = lines[i].split(',');
    
    const nameData = parseName(values[idxName]);
    const birthday = values[idxBirthday];
    
    // Calculate Age
    const age = calculateAge(birthday);

    // Format Birthday as YYYY-MM-DD
    const bParts = birthday.split('/');
    const formattedBirthday = `${bParts[2]}-${bParts[0].padStart(2, '0')}-${bParts[1].padStart(2, '0')}`;

    // Build Relatives array
    const relatives: any[] = [];

    // Helper to add relative if exists
    const addRelative = (name: string, relationship: string) => {
      const parts = parseRelativeName(name);
      if (parts.length > 0) {
        relatives.push({
          FirstName: parts[0],
          LastName: parts.slice(1).join(' '),
          Relationship: relationship
        });
      }
    };

    addRelative(values[idxFather], 'Father');
    addRelative(values[idxMother], 'Mother');
    addRelative(values[idxBrother], 'Brother');
    addRelative(values[idxSister], 'Sister');

    result.push({
      FirstName: nameData.firstName,
      LastName: nameData.lastName,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    });
  }

  console.log(JSON.stringify(result, null, 2));
}

main();