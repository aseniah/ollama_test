import * as fs from 'fs';

function parseDate(dateStr: string): Date {
  // Format: M/D/YYYY
  const parts = dateStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

function calculateAge(birthday: Date, died: Date | null): number {
  const asOf = new Date(2025, 6, 1); // July 1, 2025 (0-indexed month)
  
  let referenceDate: Date;
  if (died) {
    referenceDate = died;
  } else {
    referenceDate = asOf;
  }
  
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthday.getMonth();
  const dayDiff = referenceDate.getDate() - birthday.getDate();
  
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }
  
  return age;
}

function main() {
  const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
  const lines = csvContent.trim().split('\n');
  const header = lines[0].split(',');
  
  const result: any[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;
    
    const fields = line.split(',');
    
    // Name
    const fullName = fields[0];
    const nameParts = fullName.split(' ');
    let firstName: string;
    let lastName: string;
    
    if (nameParts.length === 2) {
      firstName = nameParts[0];
      lastName = nameParts[1];
    } else {
      // Assume first name is first word, last name is last word
      // For "John Winston Lennon" -> John Lennon? 
      // Looking at expected: John Lennon (from "John Winston Lennon")
      // So we need to drop the middle name
      if (nameParts.length >= 2) {
        firstName = nameParts[0];
        lastName = nameParts[nameParts.length - 1];
      } else {
        firstName = fullName;
        lastName = '';
      }
    }
    
    // Birthday
    const birthday = parseDate(fields[1]);
    
    // Died
    let died: Date | null = null;
    if (fields[2] !== 'null' && fields[2].trim() !== '') {
      died = parseDate(fields[2]);
    }
    
    const age = calculateAge(birthday, died);
    
    // Relatives
    const relatives: any[] = [];
    
    // Father (index 3)
    if (fields[3] !== 'null' && fields[3].trim() !== '') {
      const famParts = fields[3].split(' ');
      relatives.push({
        FirstName: famParts[0],
        LastName: famParts.slice(1).join(' '),
        Relationship: 'Father'
      });
    }
    
    // Mother (index 4)
    if (fields[4] !== 'null' && fields[4].trim() !== '') {
      const famParts = fields[4].split(' ');
      relatives.push({
        FirstName: famParts[0],
        LastName: famParts.slice(1).join(' '),
        Relationship: 'Mother'
      });
    }
    
    // Brother (index 5)
    if (fields[5] !== 'null' && fields[5].trim() !== '') {
      const famParts = fields[5].split(' ');
      relatives.push({
        FirstName: famParts[0],
        LastName: famParts.slice(1).join(' '),
        Relationship: 'Brother'
      });
    }
    
    // Sister (index 6)
    if (fields[6] !== 'null' && fields[6].trim() !== '') {
      const famParts = fields[6].split(' ');
      relatives.push({
        FirstName: famParts[0],
        LastName: famParts.slice(1).join(' '),
        Relationship: 'Sister'
      });
    }
    
    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: fields[1], // Convert to YYYY-MM-DD? Let's check expected format
      Age: age,
      Relatives: relatives
    });
  }
  
  // Check expected format: Birthday is "1940-10-09" (ISO format)
  // Need to convert from M/D/YYYY to YYYY-MM-DD
  
  for (const person of result) {
    const bd = parseDate(person.Birthday as string);
    const yyyy = bd.getFullYear();
    const mm = String(bd.getMonth() + 1).padStart(2, '0');
    const dd = String(bd.getDate()).padStart(2, '0');
    person.Birthday = `${yyyy}-${mm}-${dd}`;
  }
  
  console.log(JSON.stringify(result, null, 2));
}

main();