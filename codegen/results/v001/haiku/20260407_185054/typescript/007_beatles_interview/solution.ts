import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(process.cwd(), 'input', 'input.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');
const lines = csvContent.trim().split('\n');

const headers = lines[0].split('\t');
const nameIdx = headers.indexOf('Name');
const birthdayIdx = headers.indexOf('Birthday');
const diedIdx = headers.indexOf('Died');
const fatherIdx = headers.indexOf('Father');
const motherIdx = headers.indexOf('Mother');
const brotherIdx = headers.indexOf('Brother');
const sisterIdx = headers.indexOf('Sister');

function parseDate(dateStr: string): Date {
  const [month, day, year] = dateStr.split('/').map(Number);
  return new Date(year, month - 1, day);
}

function calculateAge(birthDate: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function extractName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  const lastName = parts[parts.length - 1];
  const firstName = parts.slice(0, -1).join(' ');
  return { firstName, lastName };
}

const referenceDate = new Date(2025, 6, 1);

const result = lines.slice(1).map((line) => {
  const fields = line.split('\t');
  const fullName = fields[nameIdx];
  const { firstName, lastName } = extractName(fullName);
  const birthdayStr = fields[birthdayIdx];
  const diedStr = fields[diedIdx];
  const fatherStr = fields[fatherIdx];
  const motherStr = fields[motherIdx];
  const brotherStr = fields[brotherIdx];
  const sisterStr = fields[sisterIdx];

  const birthDate = parseDate(birthdayStr);
  const birthDateStr = birthDate.toISOString().split('T')[0];

  const relatives: any[] = [];

  if (fatherStr && fatherStr !== 'null') {
    const { firstName: fFN, lastName: fLN } = extractName(fatherStr);
    relatives.push({
      FirstName: fFN,
      LastName: fLN,
      Relationship: 'Father'
    });
  }

  if (motherStr && motherStr !== 'null') {
    const { firstName: mFN, lastName: mLN } = extractName(motherStr);
    relatives.push({
      FirstName: mFN,
      LastName: mLN,
      Relationship: 'Mother'
    });
  }

  if (brotherStr && brotherStr !== 'null') {
    const { firstName: bFN, lastName: bLN } = extractName(brotherStr);
    relatives.push({
      FirstName: bFN,
      LastName: bLN,
      Relationship: 'Brother'
    });
  }

  if (sisterStr && sisterStr !== 'null') {
    const { firstName: sFN, lastName: sLN } = extractName(sisterStr);
    relatives.push({
      FirstName: sFN,
      LastName: sLN,
      Relationship: 'Sister'
    });
  }

  const age = calculateAge(birthDate, referenceDate);

  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthDateStr,
    Age: age,
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));
