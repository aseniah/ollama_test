import { readFileSync } from 'fs';

function calculateAge(birthDate: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function parseDate(mdy: string): Date {
  const [m, d, y] = mdy.split('/').map(Number);
  return new Date(y, m - 1, d);
}

function formatDate(d: Date): string {
  const year = d.getFullYear();
  const month = String(d.getMonth() + 1).padStart(2, '0');
  const day = String(d.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function main() {
  const csvContent = readFileSync('input/input.csv', 'utf-8').trim();
  const lines = csvContent.split('\n');
  const headers = lines[0].split(',');
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025

  const results: any[] = [];

  for (let i = 1; i < lines.length; i++) {
    const fields = lines[i].split(',');
    const row: Record<string, string> = {};
    headers.forEach((h, idx) => {
      row[h.trim()] = (fields[idx] || '').trim();
    });

    // Parse name into FirstName and LastName
    const nameParts = row['Name'].split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];

    // Parse birthday
    const birthDate = parseDate(row['Birthday']);
    const birthdayStr = formatDate(birthDate);

    // Calculate age
    let age: number;
    if (row['Died'] && row['Died'] !== 'null') {
      const deathDate = parseDate(row['Died']);
      age = calculateAge(birthDate, deathDate);
    } else {
      age = calculateAge(birthDate, referenceDate);
    }

    // Build relatives
    const relatives: any[] = [];
    const relativeFields: Array<[string, string]> = [
      ['Father', 'Father'],
      ['Mother', 'Mother'],
      ['Brother', 'Brother'],
      ['Sister', 'Sister']
    ];

    for (const [col, relationship] of relativeFields) {
      const value = row[col];
      if (value && value !== 'null') {
        const relNameParts = value.split(' ');
        const relFirstName = relNameParts[0];
        const relLastName = relNameParts[relNameParts.length - 1];
        relatives.push({
          FirstName: relFirstName,
          LastName: relLastName,
          Relationship: relationship
        });
      }
    }

    results.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayStr,
      Age: age,
      Relatives: relatives
    });
  }

  process.stdout.write(JSON.stringify(results, null, 2));
}

main();