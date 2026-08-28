import * as fs from 'fs';
import * as path from 'path';

const REFERENCE_DATE = new Date('2025-07-01');

function parseDate(dateString: string): Date | null {
  if (!dateString || dateString === 'null') return null;
  
  // Try parsing different date formats
  // CSV uses M/D/YYYY format (e.g., 10/9/1940)
  const parts = dateString.split('/');
  if (parts.length !== 3) return null;
  
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  
  // Adjust month for JS Date (0-indexed)
  const date = new Date(year, month - 1, day);
  
  if (isNaN(date.getTime())) return null;
  return date;
}

function calculateAge(birthday: Date): number {
  let age = REFERENCE_DATE.getFullYear() - birthday.getFullYear();
  const m = REFERENCE_DATE.getMonth() - birthday.getMonth();
  
  if (m < 0 || (m === 0 && REFERENCE_DATE.getDate() < birthday.getDate())) {
    age--;
  }
  
  return age;
}

function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(' ');
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  
  // Assume last part is the last name
  const lastName = parts.pop()!;
  const firstName = parts.join(' ');
  
  return { firstName, lastName };
}

function buildRelative(
  fullName: string | null, 
  relationship: string
): { FirstName: string; LastName: string; Relationship: string } | null {
  if (!fullName || fullName === 'null') return null;
  
  const { firstName, lastName } = parseName(fullName);
  if (firstName.trim() === '') return null;
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Relationship: relationship
  };
}

function processCSV(): void {
  const inputPath = path.join(__dirname, '..', 'input', 'input.csv');
  
  try {
    const csvContent = fs.readFileSync(inputPath, 'utf8');
    const lines = csvContent.split('\n').filter(line => line.trim());
    
    if (lines.length === 0) {
      console.log('[]');
      return;
    }
    
    // Skip header row
    const dataLines = lines.slice(1);
    const result: any[] = [];
    
    for (const line of dataLines) {
      const values = line.split(',');
      
      if (values.length < 4) continue; // Minimum required fields
      
      const name = values[0].trim();
      const birthdayStr = values[1].trim();
      const diedStr = values[2] ? values[2].trim() : 'null';
      const fatherName = values[3] ? values[3].trim() : 'null';
      const motherName = values[4] ? values[4].trim() : 'null';
      const brotherName = values[5] ? values[5].trim() : 'null';
      const sisterName = values[6] ? values[6].trim() : 'null';
      
      // Parse birthday and calculate age
      const birthday = parseDate(birthdayStr);
      if (!birthday) continue;
      
      const { firstName, lastName } = parseName(name);
      const age = calculateAge(birthday);
      
      // Build relatives array
      const relatives: any[] = [];
      
      const fatherRel = buildRelative(fatherName, 'Father');
      const motherRel = buildRelative(motherName, 'Mother');
      const brotherRel = buildRelative(brotherName, 'Brother');
      const sisterRel = buildRelative(sisterName, 'Sister');
      
      if (fatherRel) relatives.push(fatherRel);
      if (motherRel) relatives.push(motherRel);
      if (brotherRel) relatives.push(brotherRel);
      if (sisterRel) relatives.push(sisterRel);
      
      // Format birthday as YYYY-MM-DD
      const formattedBirthday = `${birthday.getFullYear()}-${String(birthday.getMonth() + 1).padStart(2, '0')}-${String(birthday.getDate()).padStart(2, '0')}`;
      
      result.push({
        FirstName: firstName,
        LastName: lastName,
        Birthday: formattedBirthday,
        Age: age,
        Relatives: relatives
      });
    }
    
    console.log(JSON.stringify(result, null, 3));
    
  } catch (error) {
    // If there's an error reading the file, output empty array
    console.log('[]');
  }
}

processCSV();