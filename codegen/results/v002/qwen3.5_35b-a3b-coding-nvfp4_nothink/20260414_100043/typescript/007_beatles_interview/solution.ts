import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'input.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');

const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');

const referenceDate = new Date('2025-07-01T00:00:00Z');

function parseBirthday(birthdayStr: string): Date {
  // Format: M/D/YYYY or MM/D/YYYY or M/DD/YYYY or MM/DD/YYYY
  const [month, day, year] = birthdayStr.split('/').map(Number);
  return new Date(year, month - 1, day);
}

function calculateAge(birthday: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const bMonth = birthday.getMonth();
  const bDay = birthday.getDate();
  const rMonth = referenceDate.getMonth();
  const rDay = referenceDate.getDate();

  if (rMonth < bMonth || (rMonth === bMonth && rDay < bDay)) {
    age--;
  }
  return age;
}

const result: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const parts = lines[i].split(',');
  if (parts.length < 2) continue;

  const nameParts = parts[0].trim().split(' ');
  const firstName = nameParts[0];
  const lastName = nameParts.slice(1).join(' ');

  const birthdayStr = parts[1].trim();
  const birthday = parseBirthday(birthdayStr);
  const age = calculateAge(birthday);

  const relatives: any[] = [];

  const fatherName = parts[3]?.trim();
  if (fatherName && fatherName !== 'null') {
    const fParts = fatherName.split(' ');
    relatives.push({
      FirstName: fParts[0],
      LastName: fParts.slice(1).join(' '),
      Relationship: 'Father'
    });
  }

  const motherName = parts[4]?.trim();
  if (motherName && motherName !== 'null') {
    const mParts = motherName.split(' ');
    relatives.push({
      FirstName: mParts[0],
      LastName: mParts.slice(1).join(' '),
      Relationship: 'Mother'
    });
  }

  const brotherName = parts[5]?.trim();
  if (brotherName && brotherName !== 'null') {
    const bParts = brotherName.split(' ');
    relatives.push({
      FirstName: bParts[0],
      LastName: bParts.slice(1).join(' '),
      Relationship: 'Brother'
    });
  }

  const sisterName = parts[6]?.trim();
  if (sisterName && sisterName !== 'null') {
    const sParts = sisterName.split(' ');
    relatives.push({
      FirstName: sParts[0],
      LastName: sParts.slice(1).join(' '),
      Relationship: 'Sister'
    });
  }

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));