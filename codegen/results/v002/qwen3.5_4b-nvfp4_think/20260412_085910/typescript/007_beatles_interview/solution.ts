import { readFileSync } from 'fs';

const inputFile = process.argv[2] || 'input/input.csv';
const dateFormatFile = process.argv[3] || 'input/expected_format.json';

function parseDateMMDDYYYY(dateStr: string): Date {
  const parts = dateStr.split('/');
  const month = parseInt(parts[0]);
  const day = parseInt(parts[1]);
  const year = parseInt(parts[2]);
  return new Date(year, month - 1, day);
}

function calculateAge(birthDate: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const refMonth = referenceDate.getMonth();
  const birthMonth = birthDate.getMonth();
  
  if (refMonth < birthMonth) {
    age--;
  } else if (refMonth === birthMonth && refDate.getDay() < birthDate.getDate()) {
    age--;
  }
  
  // Calculate day of month comparison properly
  const refDay = referenceDate.getDate();
  const birthDay = birthDate.getDate();
  
  if (refMonth === birthMonth) {
    if (refDay < birthDay) {
      age--;
    }
  }
  
  return Math.max(0, age);
}

function parseBirthday(dateStr: string): Date {
  const parts = dateStr.split('/');
  const month = parseInt(parts[0]);
  const day = parseInt(parts[1]);
  const year = parseInt(parts[2]);
  return new Date(year, month - 1, day);
}

function parseDied(dateStr: string): Date {
  const parts = dateStr.split('/');
  const year = parseInt(parts[0]);
  const month = parseInt(parts[1]) ?? 0;
  const day = parseInt(parts[2]) ?? 0;
  return new Date(year, month - 1, day);
}

function toISODate(date: Date): string {
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`;
}

const referenceDate = new Date('2025-07-01');

function parseCSV(csv: string): any[] {
  return csv
    .split('\n')
    .map((line) => line.trim())
    .filter((line) => line.length > 0)
    .map((line) => line.split(','))
    .filter((_, i) => i < (csv === null ? 1 : 4))
    .map((headers, i) => {
      const headerName = headers[i];
      return headers.filter((h) => h.trim() !== '');
    });
}

const csvContent = readFileSync(inputFile, 'utf-8');
const relativesData: any[] = [];

const rows = csvContent
  .split('\n')
  .filter((line) => line.length > 0 && !line.startsWith('Name'))
  .map((line) => line.split(','))
  .filter((_, i) => i < 4);

function buildRelatives(): any[] {
  const relativesMap: Map<string, string[]> = new Map();
  
  for (const row of rows) {
    if (!row || row.length < 6) continue;
    
    const fullNames = row[0] || '';
    const birthdayParts = row[1]?.split('/') || [];
    const diedParts = row[2]?.split('/') || [];
    const fatherParts = row[3]?.split('/') || [];
    const motherParts = row[4]?.split('/') || [];
    const brotherName = row[5] || '';
    const sisterName = row[6] || '';
    
    const [first, last] = fullNames.split(' ');
    
    let fatherName = '';
    let fatherRel = 'Father';
    
    if (fatherParts.length >= 3) {
      const fparts = fatherParts.map((p) => p.trim()).filter((p) => p !== '').join('-');
      if (fparts) {
        fatherName = fparts;
      } else {
        for (const name of fullNames.split(' ')) {
          const fullName = name || '';
          fatherName = fullName || '';
        }
      }
    }
    
    let motherName = '';
    let motherRel = 'Mother';
    
    if (motherParts.length >= 3) {
      const mparts = motherParts.map((p) => p.trim()).filter((p) => p !== '').join('-');
      if (mparts) {
        motherName = mparts;
      } else {
        for (const name of fullNames.split(' ')) {
          const fullName = name || '';
          motherName = fullName || '';
        }
      }
    }
    
    let brotherRel: string = 'Brother';
    if (brotherName.trim() !== '') {
      const parts = brotherName.split(' ');
      brotherRel = parts.length > 1 ? parts.slice(-1).join(', ') : parts[0] || 'Brother';
    } else if (!['null', '', 'NULL', 'N/A'].includes(brotherName)) {
      for (const name of fullNames.split(' ')) {
        const fullName = name || '';
        brotherRel = fullName ? 'Brother' : 'Unknown';
      }
    }
    
    let sisterRel: string = 'Sister';
    if (sisterName.trim() !== '') {
      const parts = sisterName.split(' ');
      sisterRel = parts.length > 1 ? parts.slice(-1).join(', ') : parts[0] || 'Sister';
    } else if (!['null', '', 'NULL', 'N/A'].includes(sisterName)) {
      for (const name of fullNames.split(' ')) {
        const fullName = name || '';
        sisterRel = fullName ? 'Sister' : 'Unknown';
      }
    }
    
    relativesMap.set(fullName, [fatherName, motherName, brotherName]);
  }
  
  return Array.from(relativesMap.entries());
}

const beatles: any[] = [];
const result: any[] = [];

for (const row of rows) {
  if (!row || row.length < 6) continue;
  
  const fullNames = row[0].trim();
  const [first, last] = fullNames.split(' ');
  
  let birthdayStr = '2025-07-01';
  if (row[1]) {
    try {
      birthdayStr = toISODate(parseBirthday(row[1]));
    } catch (e) {}
  }
  
  const birthDate = parseBirthday(birthdayStr);
  
  let age = calculateAge(birthDate, referenceDate);
  
  let relatives: any[] = [];
  
  for (const relativeEntry of buildRelatives()) {
    if (!relativeEntry || relativeEntry[0].trim() === '') continue;
    
    const [fName, mName, bName] = relativeEntry;
    if (!fName && !mName && !bName) continue;
    
    relatives.push({
      FirstName: fName?.split(' ')[0] || fName?.trim().split(' ')[0] || 'Unknown',
      LastName: fName?.split(' ')[1] || mName?.split(' ')[1] || bName?.split(' ')[1] || '',
      Relationship: relativeEntry[2]?.trim() || 'Unknown'
    });
  }
  
  result.push({
    FirstName: first.trim(),
    LastName: last.trim(),
    Birthday: toISODate(birthDate),
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));