import * as fs from 'fs';

const REF_DATE = new Date('2025-07-01T00:00:00Z');

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface OutputPerson {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

function parseDate(dateStr: string): Date | null {
  if (dateStr.toLowerCase() === 'null') return null;
  // Format: M/D/YYYY (e.g., 10/9/1940)
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;
  
  const month = parseInt(parts[0], 10) - 1; // 0-indexed
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  
  return new Date(year, month, day);
}

function calculateAge(birthday: Date, refDate: Date): number {
  let age = refDate.getFullYear() - birthday.getFullYear();
  const m = refDate.getMonth() - birthday.getMonth();
  
  if (m < 0 || (m === 0 && refDate.getDate() < birthday.getDate())) {
    age--;
  }
  
  return age;
}

function splitName(name: string): { first: string; last: string } {
  const parts = name.trim().split(/\s+/);
  // Assuming format: "FirstName LastName" or "FirstName MiddleName LastName"
  // We treat the last part as LastName and everything before as FirstName
  // Based on expected output: "John Winston Lennon" -> "John" (FirstName) "Lennon" (LastName)
  // It seems they take the first word as FirstName and the last word as LastName.
  if (parts.length === 2) {
    return { first: parts[0], last: parts[1] };
  } else {
    return { first: parts[0], last: parts[parts.length - 1] };
  }
}

function parseName(name: string): { first: string; last: string } | null {
  if (!name || name.toLowerCase() === 'null') return null;
  return splitName(name);
}

function main() {
  const csvPath = 'input/input.csv';
  
  if (!fs.existsSync(csvPath)) {
    process.stdout.write('[]');
    return;
  }

  const rawData = fs.readFileSync(csvPath, 'utf-8');
  const lines = rawData.split('\n').filter(line => line.trim() !== '');
  
  if (lines.length === 0) {
    process.stdout.write('[]');
    return;
  }

  const headers = lines[0].split(',').map(h => h.trim());
  const result: OutputPerson[] = [];

  for (let i = 1; i < lines.length; i++) {
    const row = lines[i].split(',');
    
    // Map header indices
    const idx = (field: string) => headers.indexOf(field);
    
    const name = row[idx('Name')];
    const birthdayStr = row[idx('Birthday')];
    const father = row[idx('Father')];
    const mother = row[idx('Mother')];
    const brother = row[idx('Brother')];
    const sister = row[idx('Sister')];

    const personName = parseName(name);
    if (!personName) continue;

    const birthday = parseDate(birthdayStr);
    if (!birthday) continue;

    // Format Birthday as YYYY-MM-DD
    const year = birthday.getFullYear();
    const month = String(birthday.getMonth() + 1).padStart(2, '0');
    const day = String(birthday.getDate()).padStart(2, '0');
    const formattedBirthday = `${year}-${month}-${day}`;

    const age = calculateAge(birthday, REF_DATE);

    const relatives: Relative[] = [];

    const addRelative = (nameStr: string, relationship: string) => {
      const parsed = parseName(nameStr);
      if (parsed) {
        relatives.push({
          FirstName: parsed.first,
          LastName: parsed.last,
          Relationship: relationship
        });
      }
    };

    addRelative(father, 'Father');
    addRelative(mother, 'Mother');
    addRelative(brother, 'Brother');
    addRelative(sister, 'Sister');

    result.push({
      FirstName: personName.first,
      LastName: personName.last,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    });
  }

  process.stdout.write(JSON.stringify(result, null, 2));
}

main();