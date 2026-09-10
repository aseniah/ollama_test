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

function parseDate(dateStr: string): Date {
  // Input format is M/D/YYYY or MM/DD/YYYY
  const [month, day, year] = dateStr.split('/').map(Number);
  return new Date(year, month - 1, day);
}

function calculateAge(birthdayStr: string, deathDateStr: string | null, targetDate: Date): number {
  const birthDate = parseDate(birthdayStr);
  const deathDate = deathDateStr ? parseDate(deathDateStr) : null;
  
  // The effective date to calculate age is the earlier of the targetDate or the date of death
  const endDate = (deathDate && deathDate < targetDate) ? deathDate : targetDate;
  
  let age = endDate.getFullYear() - birthDate.getFullYear();
  const m = endDate.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && endDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function splitName(fullName: string): { first: string; last: string } {
  if (!fullName || fullName === 'null') return { first: '', last: '' };
  const parts = fullName.trim().split(' ');
  if (parts.length === 1) return { first: parts[0], last: '' };
  
  // Based on the expected format, "John Winston Lennon" becomes FirstName: "John", LastName: "Lennon"
  // It seems the first word is FirstName and the last word is LastName.
  return {
    first: parts[0],
    last: parts[parts.length - 1]
  };
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf8');
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  
  const headers = lines[0].split(',');
  const rows = lines.slice(1);
  const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

  const result: Person[] = rows.map(row => {
    const values = row.split(',');
    const data: any = {};
    headers.forEach((header, index) => {
      data[header] = values[index];
    });

    const nameParts = splitName(data.Name);
    const birthday = data.Birthday;
    const died = data.Died === 'null' ? null : data.Died;

    const relatives: Relative[] = [];
    const relationshipFields = ['Father', 'Mother', 'Brother', 'Sister'];
    
    relationshipFields.forEach(relType => {
      const relName = data[relType];
      if (relName && relName !== 'null') {
        const p = splitName(relName);
        relatives.push({
          FirstName: p.first,
          LastName: p.last,
          Relationship: relType
        });
      }
    });

    return {
      FirstName: nameParts.first,
      LastName: nameParts.last,
      Birthday: birthday.split('/').reverse().join('-').replace(/^(\d{1,2})-(\d{1,2})-(\d{4})$/, (m, d, m1, y) => {
        // Convert M/D/YYYY to YYYY-MM-DD
        // The input csv is M/D/YYYY. We need to handle padding.
        const parts = birthday.split('/');
        const mm = parts[0].padStart(2, '0');
        const dd = parts[1].padStart(2, '0');
        return `${parts[2]}-${mm}-${dd}`;
      }),
      Age: calculateAge(birthday, died, targetDate),
      Relatives: relatives
    };
  });

  // Correcting the Birthday formatting logic slightly to be cleaner
  const finalResult = result.map(p => {
    const parts = p.Birthday.split('-'); // this is wrong because I just set it above
    // Let's re-process the birthday strings from the original data to be safe
    return p;
  });

  // Re-mapping birthday to ensure YYYY-MM-DD format exactly
  const sanitizedResult = result.map(p => {
    // The previous Birthday logic was a bit messy inside the map, let's just fix it here if needed
    // Actually, let's redo the birthday mapping inside the original loop.
    return p;
  });

  // Re-calculating the birthday format properly
  const formattedResult = rows.map(row => {
    const values = row.split(',');
    const data: any = {};
    headers.forEach((header, index) => { data[header] = values[index]; });
    
    const bParts = data.Birthday.split('/');
    const birthdayISO = `${bParts[2]}-${bParts[0].padStart(2, '0')}-${bParts[1].padStart(2, '0')}`;
    
    const nameParts = splitName(data.Name);
    const died = data.Died === 'null' ? null : data.Died;

    const relatives: Relative[] = [];
    ['Father', 'Mother', 'Brother', 'Sister'].forEach(relType => {
      const relName = data[relType];
      if (relName && relName !== 'null') {
        const p = splitName(relName);
        relatives.push({ FirstName: p.first, LastName: p.last, Relationship: relType });
      }
    });

    return {
      FirstName: nameParts.first,
      LastName: nameParts.last,
      Birthday: birthdayISO,
      Age: calculateAge(data.Birthday, died, targetDate),
      Relatives: relatives
    };
  });

  process.stdout.write(JSON.stringify(formattedResult, null, 2) + '\n');
}

main();