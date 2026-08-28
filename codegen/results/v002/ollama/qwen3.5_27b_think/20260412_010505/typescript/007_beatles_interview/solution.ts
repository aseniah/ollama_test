import * as fs from 'fs';

function parseCSV(csv: string): string[][] {
  const lines = csv.trim().split('\n');
  const result: string[][] = [];
  
  for (const line of lines) {
    const values: string[] = [];
    let current = '';
    let inQuotes = false;
    
    for (const char of line) {
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        values.push(current.trim());
        current = '';
      } else {
        current += char;
      }
    }
    values.push(current.trim());
    result.push(values);
  }
  
  return result;
}

function parseDate(dateStr: string): Date {
  const parts = dateStr.split('/');
  const month = parseInt(parts[0], 10) - 1;
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month, day);
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const bMonthDay = birthday.getMonth() * 30 + birthday.getDate();
  const rMonthDay = referenceDate.getMonth() * 30 + referenceDate.getDate();
  
  if (bMonthDay > rMonthDay) {
    age--;
  }
  return age;
}

function splitName(name: string): { firstName: string; lastName: string } {
  const parts = name.split(/\s+/);
  return {
    firstName: parts[0],
    lastName: parts.slice(1).join(' ')
  };
}

const input = fs.readFileSync('input/input.csv', 'utf-8');
const csvData = parseCSV(input);
const rows = csvData.slice(1);
const referenceDate = new Date(2025, 6, 1);

const result = rows.map(row => {
  const name = row[0];
  const birthday = row[1];
  const father = row[3];
  const mother = row[4];
  const brother = row[5];
  const sister = row[6];
  
  const { firstName, lastName } = splitName(name);
  const birthdayDate = parseDate(birthday);
  const age = calculateAge(birthdayDate, referenceDate);
  
  const relatives: { FirstName: string; LastName: string; Relationship: string }[] = [];
  
  if (father !== 'null') {
    const parts = father.split(/\s+/);
    relatives.push({
      FirstName: parts[0],
      LastName: parts.slice(1).join(' '),
      Relationship: 'Father'
    });
  }
  if (mother !== 'null') {
    const parts = mother.split(/\s+/);
    relatives.push({
      FirstName: parts[0],
      LastName: parts.slice(1).join(' '),
      Relationship: 'Mother'
    });
  }
  if (brother !== 'null') {
    const parts = brother.split(/\s+/);
    relatives.push({
      FirstName: parts[0],
      LastName: parts.slice(1).join(' '),
      Relationship: 'Brother'
    });
  }
  if (sister !== 'null') {
    const parts = sister.split(/\s+/);
    relatives.push({
      FirstName: parts[0],
      LastName: parts.slice(1).join(' '),
      Relationship: 'Sister'
    });
  }
  
  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: formatDate(birthdayDate),
    Age: age,
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));