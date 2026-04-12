import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a CSV file containing information about people,
 * and transforms it into a specific JSON format based on rules
 * inferred from a provided expected format.
 *
 * Transformation Rules:
 * 1. Name: Split into FirstName (first part) and LastName (last part).
 * 2. Birthday: Convert M/D/YYYY to YYYY-MM-DD.
 * 3. Age: 
 *    - If Died is provided, Age = Age at time of death.
 *    - If Died is null, Age = Age as of July 1, 2025.
 * 4. Relatives: Extract Father, Mother, Brother, Sister columns.
 *    - Skip null/empty columns.
 *    - For each relative, split name into FirstName and LastName.
 *    - Relationship is "Father", "Mother", "Brother", or "Sister".
 */

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

function parseCSVDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null' || dateStr.trim() === '') {
    return null;
  }
  // Expected format in CSV: M/D/YYYY
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  // JavaScript months are 0-indexed
  return new Date(year, month - 1, day);
}

function formatDateToYYYYMMDD(date: Date): string {
  const y = date.getFullYear();
  const m = String(date.getMonth() + 1).padStart(2, '0');
  const d = String(date.getDate()).padStart(2, '0');
  return `${y}-${m}-${d}`;
}

function calculateAge(birthDate: Date, endDate: Date): number {
  let age = endDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = endDate.getMonth() - birthDate.getMonth();
  const dayDiff = endDate.getDate() - birthDate.getDate();

  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }
  return age;
}

function main() {
  const inputFilePath = path.join('input', 'input.csv');
  
  let csvContent: string;
  try {
    csvContent = fs.readFileSync(inputFilePath, 'utf-8');
  } catch (err) {
    // If file not found, exit gracefully
    process.exit(0);
  }

  const lines = csvContent.split(/\r?\n/).map(line => line.trim()).filter(line => line.length > 0);
  if (lines.length < 2) {
    console.log('[]');
    return;
  }

  // The first line is the header: Name,Birthday,Died,Father,Mother,Brother,Sister
  const rows = lines.slice(1);
  const refDate = new Date(2025, 6, 1); // July 1, 2025 (Month index 6)

  const result: Person[] = rows.map(row => {
    const columns = row.split(',');
    
    // 1. Process Name
    const fullName = columns[0];
    const nameParts = fullName.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];

    // 2. Process Birthday
    const birthDate = parseCSVDate(columns[1]);
    if (!birthDate) {
      throw new Error(`Invalid birthday in row: ${row}`);
    }
    const birthdayStr = formatDateToYYYYMMDD(birthDate);

    // 3. Process Death and Age
    const diedDate = parseCSVDate(columns[2]);
    const endDate = diedDate !== null ? diedDate : refDate;
    const age = calculateAge(birthDate, endDate);

    // 4. Process Relatives
    const relatives: Relative[] = [];
    const relTypes = ['Father', 'Mother', 'Brother', 'Sister'];
    const relIndices = [3, 4, 5, 6];

    for (let i = 0; i < relTypes.length; i++) {
      const relName = columns[relIndices[i]];
      if (relName && relName !== 'null' && relName.trim() !== '') {
        const relParts = relName.trim().split(' ');
        relatives.push({
          FirstName: relParts[0],
          LastName: relParts[relParts.length - 1],
          Relationship: relTypes[i]
        });
      }
    }

    return {
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayStr,
      Age: age,
      Relatives: relatives
    };
  });

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();