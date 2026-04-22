import { readFileSync } from 'fs';
import { join } from 'path';

const csvPath = join(process.cwd(), 'input', 'input.csv');
const csvData = readFileSync(csvPath, 'utf-8');
const lines = csvData.trim().split('\n');
const dataLines = lines.slice(1);

const result: any[] = [];
const refDate = new Date(2025, 6, 1); // July 1, 2025

for (const line of dataLines) {
  const cols = line.split(',');
  const name = cols[0].trim();
  const birthdayStr = cols[1].trim();
  const father = cols[3].trim();
  const mother = cols[4].trim();
  const brother = cols[5].trim();
  const sister = cols[6].trim();

  const nameParts = name.split(' ');
  const firstName = nameParts[0];
  const lastName = nameParts[nameParts.length - 1];

  const [monthStr, dayStr, yearStr] = birthdayStr.split('/');
  const month = parseInt(monthStr, 10);
  const day = parseInt(dayStr, 10);
  const year = parseInt(yearStr, 10);
  
  const birthdayDate = new Date(year, month - 1, day);
  
  let age = refDate.getFullYear() - birthdayDate.getFullYear();
  if (refDate.getMonth() < birthdayDate.getMonth() || 
      (refDate.getMonth() === birthdayDate.getMonth() && refDate.getDate() < birthdayDate.getDate())) {
    age--;
  }

  const relatives: any[] = [];
  const addRelative = (rel: string, relType: string) => {
    if (rel !== 'null') {
      const parts = rel.split(' ');
      relatives.push({ FirstName: parts[0], LastName: parts[parts.length - 1], Relationship: relType });
    }
  };

  addRelative(father, 'Father');
  addRelative(mother, 'Mother');
  addRelative(brother, 'Brother');
  addRelative(sister, 'Sister');

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`,
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));