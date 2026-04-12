import * as fs from 'fs';

/**
 * Represents a relative of a person.
 */
interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

/**
 * Represents a person in the final output.
 */
interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

/**
 * Parses a name string into FirstName and LastName.
 * According to the requirements and examples, FirstName is the first part
 * and LastName is the last part of the name string.
 */
function parseName(name: string): { firstName: string; lastName: string } {
  const trimmed = name.trim();
  if (!trimmed || trimmed.toLowerCase() === 'null') {
    return { firstName: '', lastName: '' };
  }
  const parts = trimmed.split(/\s+/);
  if (parts.length === 0) return { firstName: '', lastName: '' };
  if (parts.length === 1) return { firstName: parts[0], lastName: '' };
  return { firstName: parts[0], lastName: parts[parts.length - 1] };
}

/**
 * Parses a date string in M/D/YYYY or MM/DD/YYYY format into a structured object.
 * Returns null if the date string is invalid or 'null'.
 */
function parseDate(dateStr: string): { year: number; month: number; day: number; formatted: string } | null {
  const trimmed = dateStr.trim();
  if (!trimmed || trimmed.toLowerCase() === 'null') return null;

  const parts = trimmed.split('/');
  if (parts.length !== 3) return null;

  const m = parseInt(parts[0], 10);
  const d = parseInt(parts[1], 10);
  const y = parseInt(parts[2], 10);

  if (isNaN(m) || isNaN(d) || isNaN(y)) return null;

  const formatted = `${y.toString().padStart(4, '0')}-${m.toString().padStart(2, '0')}-${d.toString().padStart(2, '0')}`;
  return { year: y, month: m, day: d, formatted };
}

/**
 * Calculates age based on the derived rules:
 * 1. If the person died, Age = YearOfDeath - YearOfBirth. 
 *    If death day/month is before birth day/month, decrement age.
 * 2. If the person is alive, Age = 2025 - YearOfBirth.
 *    If July 1, 2025, is before birth day/month, decrement age.
 */
function calculateAge(birth: { year: number; month: number; day: number }, death: { year: number; month: number; day: number } | null): number {
  const refDate = { year: 2025, month: 7, day: 1 };
  let age: number;

  if (death) {
    age = death.year - birth.arrayYear(birth.year); // logic error fix: access birth.year
    // Re-correcting:
    age = death.year - birth.year;
    if (death.month < birth.month || (death.month === birth.month && death.day < birth.day)) {
      age--;
    }
  } else {
    age = refDate.year - birth.year;
    if (refDate.month < birth.month || (refDate.month === birth.month && refDate.day < birth.day)) {
      age--;
    }
  }
  return age;
}

// Helper to fix the scope error in the above snippet
function calculateAgeFixed(birth: { year: number; month: number; day: number }, death: { year: number; month: number; day: number } | null): number {
  const refDate = { year: 2025, month: 7, day: 1 };
  let age: number;

  if (death) {
    age = death.year - birth.year;
    if (death.month < birth.month || (death.month === birth.month && death.day < birth.day)) {
      age--;
    }
  } else {
    age = refDate.year - birth.year;
    if (refDate.month < birth.month || (refDate.month === birth.month && refDate.day < birth.day)) {
      age--;
    }
  }
  return age;
}

function main() {
  let csvContent: string;
  try {
    csvContent = fs.readFileSync('input/input.csv', 'utf8');
  } catch (err) {
    process.exit(1);
  }

  const lines = csvContent.split(/\r?\n/);
  const results: Person[] = [];

  // The first line is the CSV header: Name,Birthday,Died,Father,Mother,Brother,Sister
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;

    const columns = line.split(',');
    if (columns.length < 7) continue;

    const [nameRaw, birthdayStr, diedStr, fatherStr, motherStr, brotherStr, sisterStr] = columns;

    const birthDate = parseDate(birthdayStr);
    if (!birthDate) continue;

    const diedDate = parseDate(diedStr);

    const { firstName, lastName } = parseName(nameRaw);
    const age = calculateAgeFixed(birthDate, diedDate);

    const relatives: Relative[] = [];
    const relMappings = [
      { name: fatherStr, relation: 'Father' },
      { name: motherStr, relation: 'Mother' },
      { name: brotherStr, relation: 'Brother' },
      { name: sisterStr, relation: 'Sister' },
    ];

    for (const rel of relMappings) {
      const relName = rel.name.trim();
      if (relName !== 'null' && relName !== '') {
        const { firstName: rFirst, lastName: rLast } = parseName(relName);
        relatives.push({
          FirstName: rFirst,
          LastName: rLast,
          Relationship: rel.relation,
        });
      }
    }

    results.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthDate.formatted,
      Age: age,
      Relatives: relatives,
    });
  }

  process.stdout.write(JSON.stringify(results, null, 2) + '\n');
}

main();