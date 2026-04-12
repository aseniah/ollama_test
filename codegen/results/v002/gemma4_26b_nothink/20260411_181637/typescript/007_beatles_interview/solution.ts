import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the output JSON objects.
 */
interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
ly}

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

/**
 * Helper to parse a name into First and Last name parts.
 * Based on the input.csv, the name field contains full names.
 * The expected output shows the first part as FirstName and the rest as LastName.
 * However, looking at "John Winston Lennon", the expected output shows "John" and "Lennon".
 * Wait, checking the example: 
 * Input: "John Winston Lennon" -> Output: "John", "Lennon"
 * This implies we take the first word as FirstName and the last word as LastName.
 */
function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  const firstName = parts[0];
  const lastName = parts[parts.length - 1];
  return { firstName, lastName };
}

/**
 * Helper to parse relative names into the expected format.
 */
function parseRelative(fullName: string, relationship: string): Relative | null {
  if (!fullName || fullName.toLowerCase() === 'null') return null;
  const { firstName, lastName } = parseName(fullName);
  return {
    FirstName: firstName,
    LastName: lastName,
    Relationship: relationship
  };
}

/**
 * Calculates age as of July 1, 2025.
 */
function calculateAge(birthdayStr: string): number {
  // Input format in CSV is M/D/YYYY (e.g., 10/9/1940)
  const [month, day, year] = birthdayStr.split('/').map(Number);
  const birthday = new Date(year, month - 1, day);
  const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month index 6 is July)

  let age = targetDate.getFullYear() - birthday.getFullYear();
  const m = targetDate.getMonth() - birthday.getMonth();
  if (m < 0 || (m === 0 && targetDate.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

/**
 * Formats date to YYYY-MM-DD.
 */
function formatDate(dateStr: string): string {
  const [month, day, year] = dateStr.split('/').map(Number);
  const m = String(month).padStart(2, '0');
  const d = String(day).padStart(2, '0');
  const y = String(year);
  return `${y}-${m}-${d}`;
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'input.csv');

  if (!fs.existsSync(inputPath)) {
    process.exit(0);
  }

  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  const lines = csvContent.trim().split('\n');
  
  // Remove header
  const headers = lines[0].split(',');
  const rows = lines.slice(1);

  const result: Person[] = [];

  for (const row of rows) {
    const columns = row.split(',');
    if (columns.length < 1) continue;

    // CSV columns: Name, Birthday, Died, Father, Mother, Brother, Sister
    const fullName = columns[0];
    const birthdayRaw = columns[1];
    // columns[2] is Died (not used in output)
    const fatherRaw = columns[3];
    const motherRaw = columns[4];
    const brotherRaw = columns[5];
    const sisterRaw = columns[6];

    const { firstName, lastName } = parseName(fullName);
    
    const relatives: Relative[] = [];
    
    const fRel = parseRelative(fatherRaw, 'Father');
    if (fRel) relatives.push(fRel);
    
    const mRel = parseRelative(motherRaw, 'Mother');
    if (mRel) relatives.push(mRel);
    
    const bRel = parseRelative(brotherRaw, 'Brother');
    if (bRel) relatives.push(bRel);
    
    const sRel = parseRelative(sisterRaw, 'Sister');
    if (sRel) relatives.push(sRel);

    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: formatDate(birthdayRaw),
      Age: calculateAge(birthdayRaw),
      Relatives: relatives
    });
  }

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();