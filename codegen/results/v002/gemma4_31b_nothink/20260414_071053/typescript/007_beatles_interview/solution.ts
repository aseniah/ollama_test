import * as fs from 'fs';
import * as path from 'path';

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

/**
 * Calculates age based on a birthday and a reference date (2025-07-01).
 * If the person has a death date, the age is calculated as of the death date.
 */
function calculateAge(birthdayStr: string, deathDateStr: string | null): number {
  const referenceDate = new Date('2025-07-01');
  
  // CSV Date format is M/D/YYYY
  const parseDate = (s: string) => {
    const [m, d, y] = s.split('/').map(Number);
    return new Date(y, m - 1, d);
  };

  const birth = parseDate(birthdayStr);
  const end = deathDateStr ? parseDate(deathDateStr) : referenceDate;

  let age = end.getFullYear() - birth.getFullYear();
  const m = end.getMonth() - birth.getMonth();
  if (m < 0 || (m === 0 && end.getDate() < birth.getDate())) {
    age--;
  }
  return age;
}

/**
 * Splits a full name into First and Last names.
 * Based on the expected format: 
 * "John Winston Lennon" -> First: "John", Last: "Lennon"
 */
function splitName(fullName: string): { first: string; last: string } {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 0) return { first: '', last: '' };
  if (parts.length === 1) return { first: parts[0], last: '' };
  return {
    first: parts[0],
    last: parts[parts.length - 1]
  };
}

function main() {
  const csvPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(csvPath, 'utf-8');
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  
  const headers = lines[0].split(',');
  const dataRows = lines.slice(1);

  const result: Person[] = dataRows.map(row => {
    const values = row.split(',');
    const entry: Record<string, string> = {};
    headers.forEach((header, index) => {
      entry[header] = values[index] || '';
    });

    const { first: firstName, last: lastName } = splitName(entry['Name']);
    const birthday = entry['Birthday'];
    const died = entry['Died'] === 'null' ? null : entry['Died'];

    const relatives: Relative[] = [];
    const relativeCols = ['Father', 'Mother', 'Brother', 'Sister'];

    relativeCols.forEach(relType => {
      const relName = entry[relType];
      if (relName && relName !== 'null') {
        const { first, last } = splitName(relName);
        relatives.push({
          FirstName: first,
          LastName: last,
          Relationship: relType
        });
      }
    });

    return {
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthday.split('/').reverse().join('-').replace(/-/g, (m, i) => {
        // Simple transform for M/D/YYYY to YYYY-MM-DD
        return m; 
      }), // This is a bit messy, let's do it properly
      Age: calculateAge(birthday, died),
      Relatives: relatives
    };
  });

  // Fixing the Birthday string format manually to YYYY-MM-DD
  const formattedResult = result.map(p => {
    const rawBday = (content.split(/\r?\n/).filter(l => l.trim() !== '').slice(1).find(l => l.startsWith(p.FirstName)) || '').split(',')[1];
    // To be safe, we just need to map the CSV birth date string to ISO
    const [m, d, y] = p.Birthday.split('/').map(v => v.padStart(2, '0'));
    // Wait, the Birthday was already modified in the map. Let's just re-parse the CSV value.
    return p;
  });

  // Redoing the birthday format properly
  const finalOutput = result.map(p => {
    // Since I messed up the birthday string in the previous map, 
    // let's assume the logic is: extract date from CSV and format.
    // Re-calculating birthday string from the original logic logic.
    return p;
  });

  // Let's refine the whole loop to avoid the string mess
}

// Revised main for cleanliness
function solve() {
    const csvPath = path.join('input', 'input.csv');
    const content = fs.readFileSync(csvPath, 'utf-8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    const headers = lines[0].split(',');
    
    const output = lines.slice(1).map(row => {
        const values = row.split(',');
        const data: any = {};
        headers.forEach((h, i) => data[h] = values[i]);

        const nameParts = data['Name'].split(' ');
        const birthday = data['Birthday'];
        const died = data['Died'] === 'null' ? null : data['Died'];

        // Date format M/D/YYYY -> YYYY-MM-DD
        const [m, d, y] = birthday.split('/');
        const isoBirthday = `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}`;

        const relatives = [];
        ['Father', 'Mother', 'Brother', 'Sister'].forEach(rel => {
            if (data[rel] && data[rel] !== 'null') {
                const relParts = data[rel].split(' ');
                relatives.push({
                    FirstName: relParts[0],
                    LastName: relParts[relParts.length - 1],
                    Relationship: rel
                });
            }
        });

        return {
            FirstName: nameParts[0],
            LastName: nameParts[nameParts.length - 1],
            Birthday: isoBirthday,
            Age: calculateAge(birthday, died),
            Relatives: relatives
        };
    });

    process.stdout.write(JSON.stringify(output, null, 2) + '\n');
}

solve();