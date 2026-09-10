import * as fs from 'fs';
import * as path from 'path';

/**
 * Parses a date string in the format M/D/YYYY into a Date object.
 * We treat these as local dates to ensure consistent calculations.
 */
function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;
  const month = parseInt(parts[0], 10) - 1;
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month, day);
}

/**
 * Formats a Date object to YYYY-MM-DD string.
 */
function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

/**
 * Calculates age relative to a reference date.
 */
function calculateAge(birthDate: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const m = referenceDate.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

/**
 * Splits a full name into First and Last name based on the expected output format.
 * First name is the first word, Last name is the last word.
 */
function splitName(fullName: string) {
  const parts = fullName.trim().split(/\s+/);
  return {
    FirstName: parts[0],
    LastName: parts[parts.length - 1],
  };
}

function solve() {
  const csvPath = path.join('input', 'input.csv');
  const csvContent = fs.readFileSync(csvPath, 'utf-8');
  const lines = csvContent.split(/\r?\n/).filter(line => line.trim().length > 0);
  
  if (lines.length === 0) {
    console.log('[]');
    return;
  }

  const headers = lines[0].split(',');
  const dataLines = lines.slice(1);
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

  const result = dataLines.map(line => {
    const values = line.split(',');
    const row: Record<string, string> = {};
    headers.forEach((header, index) => {
      row[header] = values[index];
    });

    const { FirstName, LastName } = splitName(row['Name']);
    const birthday = parseDate(row['Birthday']);
    const diedDate = parseDate(row['Died']);
    
    if (!birthday) throw new Error(`Missing birthday for ${row['Name']}`);

    // Age is calculated as of the earlier of: July 1, 2025 or date of death
    const calculationDate = (diedDate && diedDate < referenceDate) ? diedDate : referenceDate;
    const age = calculateAge(birthday, calculationDate);

    const relatives = [];
    const relativeColumns = ['Father', 'Mother', 'Brother', 'Sister'];
    relativeColumns.forEach(col => {
      const relativeName = row[col];
      if (relativeName && relativeName !== 'null') {
        const { FirstName: rFirst, LastName: rLast } = splitName(relativeName);
        relatives.push({
          FirstName: rFirst,
          LastName: rLast,
          Relationship: col
        });
      }
    });

    return {
      FirstName,
      LastName,
      Birthday: formatDate(birthday),
      Age: age,
      Relatives: relatives
    };
  });

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

solve();