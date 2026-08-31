import * as fs from 'fs';

function parseCsvLine(line: string): string[] {
  const parts: string[] = [];
  let current = '';
  let inQuotes = false;
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    if (inQuotes) {
      if (ch === '"') {
        if (line[i + 1] === '"') {
          current += '"';
          i++;
        } else {
          inQuotes = false;
        }
      } else {
        current += ch;
      }
    } else {
      if (ch === '"') {
        inQuotes = true;
      } else if (ch === ',') {
        parts.push(current);
        current = '';
      } else {
        current += ch;
      }
    }
  }
  parts.push(current);
  return parts;
}

function readCsv(path: string): string[][] {
  const content = fs.readFileSync(path, 'utf-8');
  const lines = content.split('\n').filter(l => l.trim() !== '');
  return lines.map(l => parseCsvLine(l));
}

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  // Format: M/D/YYYY
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  if (isNaN(month) || isNaN(day) || isNaN(year)) return null;
  // Use UTC to avoid timezone issues
  return new Date(Date.UTC(year, month - 1, day));
}

function formatISODate(date: Date): string {
  const year = date.getUTCFullYear();
  const month = date.getUTCMonth() + 1;
  const day = date.getUTCDate();
  const mm = month.toString().padStart(2, '0');
  const dd = day.toString().padStart(2, '0');
  return `${year}-${mm}-${dd}`;
}

function calculateAge(birthday: string, asOf: Date): number {
  const bd = parseDate(birthday);
  if (!bd) return 0;
  let age = asOf.getUTCFullYear() - bd.getUTCFullYear();
  const monthDiff = asOf.getUTCMonth() - bd.getUTCMonth();
  if (monthDiff < 0 || (monthDiff === 0 && asOf.getUTCDate() < bd.getUTCDate())) {
    age--;
  }
  return age;
}

function splitFullName(name: string): { first: string; last: string } {
  const parts = name.trim().split(/\s+/);
  if (parts.length === 1) {
    return { first: parts[0], last: '' };
  }
  const last = parts[parts.length - 1];
  const first = parts.slice(0, -1).join(' ');
  return { first, last };
}

function main() {
  const rows = readCsv('input/input.csv');
  const headers = rows[0];
  const dataRows = rows.slice(1);

  // Find column indices
  const nameIdx = headers.indexOf('Name');
  const birthdayIdx = headers.indexOf('Birthday');
  const diedIdx = headers.indexOf('Died');
  const fatherIdx = headers.indexOf('Father');
  const motherIdx = headers.indexOf('Mother');
  const brotherIdx = headers.indexOf('Brother');
  const sisterIdx = headers.indexOf('Sister');

  const asOf = new Date(Date.UTC(2025, 6, 1)); // July 1, 2025 (month is 0-indexed, so 6 = July)

  const result: Person[] = [];

  for (const row of dataRows) {
    if (row.length < headers.length) continue;

    const fullName = row[nameIdx].trim();
    const { first: firstName, last: lastName } = splitFullName(fullName);

    const birthday = row[birthdayIdx].trim();
    const birthdayDate = parseDate(birthday);
    if (!birthdayDate) continue;
    const birthdayISO = formatISODate(birthdayDate);

    const age = calculateAge(birthday, asOf);

    const relatives: Relative[] = [];

    // Father
    const fatherName = row[fatherIdx].trim();
    if (fatherName && fatherName !== 'null') {
      const fParts = fatherName.split(/\s+/);
      const fFirst = fParts[0];
      const fLast = fParts.slice(1).join(' ');
      relatives.push({ FirstName: fFirst, LastName: fLast, Relationship: 'Father' });
    }

    // Mother
    const motherName = row[motherIdx].trim();
    if (motherName && motherName !== 'null') {
      const mParts = motherName.split(/\s+/);
      const mFirst = mParts[0];
      const mLast = mParts.slice(1).join(' ');
      relatives.push({ FirstName: mFirst, LastName: mLast, Relationship: 'Mother' });
    }

    // Brother
    const brotherName = row[brotherIdx].trim();
    if (brotherName && brotherName !== 'null') {
      const bParts = brotherName.split(/\s+/);
      const bFirst = bParts[0];
      const bLast = bParts.slice(1).join(' ');
      relatives.push({ FirstName: bFirst, LastName: bLast, Relationship: 'Brother' });
    }

    // Sister
    const sisterName = row[sisterIdx].trim();
    if (sisterName && sisterName !== 'null') {
      const sParts = sisterName.split(/\s+/);
      const sFirst = sParts[0];
      const sLast = sParts.slice(1).join(' ');
      relatives.push({ FirstName: sFirst, LastName: sLast, Relationship: 'Sister' });
    }

    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayISO,
      Age: age,
      Relatives: relatives
    });
  }

  console.log(JSON.stringify(result, null, 2));
}

main();