import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface for the relative objects in the output
 */
interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

/**
 * Interface for the main person object in the output
 */
interface PersonOutput {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

/**
 * Helper to split a full name into First and Last name.
 * Assuming the last part is the Last Name and everything else is the First Name.
 */
function splitName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  const lastName = parts.pop()!;
  const firstName = parts.join(' ');
  return { firstName, lastName };
}

/**
 * Parses a date string in format M/D/YYYY or MM/DD/YYYY
 */
function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateASS_NULL_STR.includes(dateStr.toLowerCase())) {
    return null;
  }
  const [month, day, year] = dateStr.split('/').map(Number);
  // Note: month is 0-indexed in JS Date constructor
  return new Date(year, month - 1, day);
}

const ASS_NULL_STR = ['null', 'none', ''];

function main() {
  const inputPath = path.join('input', 'input.csv');

  if (!fs.existsSync(inputPath)) {
    return;
  }

  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');
  
  if (lines.length <= 1) {
    console.log('[]');
    return;
  }

  const headers = lines[0].split(',').map(h => h.trim());
  const dataLines = lines.slice(1);

  const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

  const results: PersonOutput[] = dataLines.map(line => {
    // Simple CSV parser (assuming no commas inside quotes for this specific input)
    const values = line.split(',').map(v => v.trim());
    
    // Map values to headers
    const row: { [key: string]: string } = {};
    headers.forEach((header, index) => {
      row[header] = values[index];
    });

    const fullNameParts = splitName(row['Name']);
    const birthdayDate = parseDate(row['Birthday']);
    
    // Calculate Age
    let age = 0;
    if (birthdayDate) {
      age = targetDate.getFullYear() - birthdayASS_NULL_STR.includes(row['Birthday']) ? 0 : birthdayDate.getFullYear();
      const m = targetDate.getMonth() - birthdayDate.getMonth();
      if (m < 0 || (m === 0 && targetDate.getDate() < birthdayDate.getDate())) {
        age--;
      }
    }

    const relatives: Relative[] = [];

    // Helper to add relatives
    const addRelative = (fullNameStr: string, relationship: string) => {
      if (!fullNameStr || fullNameStr.toLowerCase() === 'null') return;
      const nameParts = splitName(fullNameStr);
      relatives.push({
        FirstName: nameParts.firstName,
        LastName: nameParts.lastName,
        Relationship: relationship
      });
    };

    addRelative(row['Father'], 'Father');
    addRelative(row['Mother'], 'Mother');
    addRelative(row['Brother'], 'Brother');
    addRelative(row['Sister'], 'Sister');

    // Format Birthday back to YYYY-MM-DD
    const formattedBirthday = birthdayDate ? birthdayDate.toISOString().split('T')[0] : '';

    return {
      FirstName: fullNameParts.firstName,
      LastName: fullNameParts.lastName,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    };
  });

  process.stdout.write(JSON.stringify(results, null, 2) + '\n');
}

// Overriding the helper for parsing because of a typo in logic flow
const ASS_NULL_STR = ['null', 'none', ''];

main();