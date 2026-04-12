import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface for the CSV row structure
 */
interface CSVRow {
  Name: string;
  Birthday: string;
  Died: string;
  Father: string;
  Mother: string;
  Brother: string;
  Sister: string;
}

/**
 * Interface for the Final JSON Output structure
 */
interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface PersonOutput {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

function calculateAge(birthdayStr: string, referenceDate: Date): number {
  // Handle multiple date formats (MM/DD/YYYY or YYYY-MM-DD)
  // The input CSV uses M/D/YYYY based on the sample
  const parts = birthdayStr.split('/');
  let birthDate: Date;

  if (parts.length === 3) {
    // Assume M/D/YYYY
    const [month, day, year] = parts.map(Number);
    birthDate = new Date(year, month - 1, day);
  } else {
    birthDate = new Date(birthdayStr);
  }

  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  const dayDiff = referenceDate.getDate() - birthDate.getDate();

  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }

  return age;
  }

function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(' ');
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  const lastName = parts.pop()!;
  const firstName = parts.join(' ');
  return { firstName, lastName };
}

function parseRelative(name: string, relationship: string): Relative | null {
  if (!name || name === 'null') return null;
  
  const { firstName, lastName } = parseName(name);
  return {
    FirstName: firstName,
    LastName: lastName,
    Relationship: relationship
  };
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'input.csv');
  
  if (!fs.existsSync(inputPath)) {
    return;
  }

  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');
  
  if (lines.length < 2) return;

  const headers = lines[0].split(',').map(h => h.trim());
  const rows = lines.slice(1);

  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

  const result: PersonOutput[] = rows.map(row => {
    const values = row.split(',').map(v => v.trim());
    const data: Record<string, string> = {};
    headers.forEach((h, i) => {
      data[h] = values[i];
    });

    const { firstName, lastName } = parseName(data['Name']);
    
    // Formatting Birthday to YYYY-MM-DD
    const bParts = data['Birthday'].split('/');
    const bDate = new Date(Number(bParts[2]), Number(bParts[0]) - 1, Number(bParts[1]));
    const formattedBirthday = bDate.toISOString().split('T')[0];

    const age = calculateAge(data['Birthday'], referenceDate);

    const relatives: Relative[] = [];
    
    const relMap: Record<string, string> = {
      'Father': data['Father'],
      'Mother': data['Mother'],
      'Brother': data['Brother'],
      'Sister': data['Sister']
    };

    for (const [relType, relName] of Object.entries(relMap)) {
      const relObj = parseRelative(relName, relType);
      if (relObj) {
        relatives.push(relObj);
      }
    }

    return {
      FirstName: firstName,
      LastName: lastName,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    };
  });

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();