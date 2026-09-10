import * as fs from 'fs';
import * as path from 'path';

/**
 * Helper to parse date from "M/D/YYYY" format to a Date object.
 */
function parseCsvDate(dateStr: string): Date | null {
  if (!dateStr || dateStr.toLowerCase() === 'null') return null;
  const [month, day, year] = dateStr.split('/').map(Number);
  // Date constructor: year, month (0-indexed), day
  return new Date(year, month - 1, day);
}

/**
 * Helper to format Date object to "YYYY-MM-DD".
 */
function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

/**
 * Calculates age based on a birth date and a reference date.
 */
function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const m = referenceDate.getMonth() - birthday.getMonth();
  if (m < 0 || (m === 0 && referenceDate.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

/**
 * Splits a full name into First and Last name.
 * Takes the first word as FirstName and the last word as LastName.
 */
function parseName(fullName: string) {
  if (!fullName || fullName.toLowerCase() === 'null') return null;
  const parts = fullName.trim().split(/\s+/);
  return {
    FirstName: parts[0],
    LastName: parts[parts.length - 1],
  };
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  
  const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');
  if (lines.length === 0) return;

  const headers = lines[0].split(',');
  const dataLines = lines.slice(1);

  const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

  const result = dataLines.map(line => {
    const values = line.split(',');
    const row: Record<string, string> = {};
    headers.forEach((header, index) => {
      row[header] = values[index];
    });

    const birthday = parseCsvDate(row['Birthday']);
    const died = parseCsvDate(row['Died']);
    const nameParts = parseName(row['Name']);

    if (!birthday || !nameParts) return null;

    const referenceDate = died ? died : targetDate;
    const age = calculateAge(birthday, referenceDate);

    const relatives: any[] = [];
    const relativeCols = ['Father', 'Mother', 'Brother', 'Sister'];
    relativeCols.forEach(rel => {
      const relName = row[rel];
      const parsedRel = parseName(relName);
      if (parsedRel) {
        relatives.push({
          FirstName: parsedRel.FirstName,
          LastName: parsedRel.LastName,
          Relationship: rel,
        });
      }
    });

    return {
      FirstName: nameParts.FirstName,
      LastName: nameParts.LastName,
      Birthday: formatDate(birthday),
      Age: age,
      Relatives: relatives,
    };
  }).filter(item => item !== null);

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();