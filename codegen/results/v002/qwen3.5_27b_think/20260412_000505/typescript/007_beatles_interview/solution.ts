import * as fs from 'fs';
import * as path from 'path';

interface OutputPerson {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Array<{
    FirstName: string;
    LastName: string;
    Relationship: string;
  }>;
}

function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  
  // Check if birthday has occurred by the reference date
  if (
    referenceDate.getMonth() < birthday.getMonth() ||
    (referenceDate.getMonth() === birthday.getMonth() && referenceDate.getDate() < birthday.getDate())
  ) {
    age--;
  }
  
  return age;
}

function parseDate(dateString: string): Date {
  const [month, day, year] = dateString.split('/').map(Number);
  return new Date(year, month - 1, day);
}

function formatBirthday(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function splitName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  const firstName = parts[0];
  const lastName = parts.slice(1).join(' ');
  return { firstName, lastName };
}

function main() {
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025
  
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  const lines = content.trim().split('\n');
  const header = lines[0].split(',');
  const data = lines.slice(1);
  
  const results: OutputPerson[] = [];
  
  for (const line of data) {
    const values = line.split(',');
    const rowData: Record<string, string> = {};
    
    for (let i = 0; i < header.length; i++) {
      rowData[header[i]] = values[i] || '';
    }
    
    const { firstName, lastName } = splitName(rowData['Name']);
    const birthday = parseDate(rowData['Birthday']);
    const age = calculateAge(birthday, referenceDate);
    
    const relatives: OutputPerson['Relatives'][0][] = [];
    const relativeFields = ['Father', 'Mother', 'Brother', 'Sister'];
    
    for (const field of relativeFields) {
      const value = rowData[field];
      if (value && value !== 'null') {
        const { firstName: relFirstName, lastName: relLastName } = splitName(value);
        relatives.push({
          FirstName: relFirstName,
          LastName: relLastName,
          Relationship: field
        });
      }
    }
    
    results.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: formatBirthday(birthday),
      Age: age,
      Relatives: relatives
    });
  }
  
  console.log(JSON.stringify(results, null, 2));
}

main();