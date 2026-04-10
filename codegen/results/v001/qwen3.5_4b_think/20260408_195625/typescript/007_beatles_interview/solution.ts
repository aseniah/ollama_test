import fs from 'fs';

const expectedFormat = JSON.parse(fs.readFileSync('input/expected_format.json', 'utf8'));

const referenceDate = new Date('2025-07-01');

const parseDate = (dateStr: string): Date | null => {
  if (!dateStr || !dateStr.trim()) return null;
  
  const cleanDate = dateStr.trim();
  
  // Try YYYY-MM-DD format
  const match = cleanDate.match(/(\d{4})-(\d{1,2})-(\d{1,2})/);
  if (match) {
    const year = parseInt(match[1], 10);
    const month = parseInt(match[2], 10);
    const day = parseInt(match[3], 10);
    return new Date(year, month - 1, day);
  }
  
  // Try DD/MM/YYYY format
  const match2 = cleanDate.match(/(\d{1,2})\/(\d{1,2})\/(\d{4})/);
  if (match2) {
    const day = parseInt(match2[1], 10);
    const month = parseInt(match2[2], 10);
    const year = parseInt(match2[3], 10);
    return new Date(year, month - 1, day);
  }
  
  // Try MM/DD/YYYY format
  const match3 = cleanDate.match(/(\d{1,2})\/(\d{1,2})\/(\d{4})/);
  if (match3) {
    const month = parseInt(match3[1], 10);
    const day = parseInt(match3[2], 10);
    const year = parseInt(match3[3], 10);
    return new Date(year, month - 1, day);
  }
  
  return null;
};

const calculateAge = (birthDate: Date, referenceDate: Date): number => {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const nextYear = new Date(referenceDate.getFullYear() + 1, referenceDate.getMonth(), referenceDate.getDate());
  if (nextYear <= birthDate) {
    age--;
  }
  return age;
};

const outputFormat = expectedFormat;

// Read CSV
const csvContent = fs.readFileSync('input/input.csv', 'utf8');
const lines = csvContent.trim().split('\n');

if (lines.length < 2) {
  console.log('[]');
  process.exit(0);
}

const header = lines[0];
const dataLines = lines.slice(1);

const findDateColumnIndex = (header: string): number | null => {
  for (let i = 0; i < header.length; i++) {
    if (header[i].toLowerCase().includes('date') || header[i].toLowerCase().includes('birth') || header[i].toLowerCase().includes('born')) {
      return i;
    }
  }
  return null;
};

const dateColIndex = findDateColumnIndex(header) || 0;

const outputItems: (typeof outputFormat)['items'][number] = [];

for (let i = 0; i < dataLines.length; i++) {
  const line = dataLines[i];
  const dateStr = line.split(',')[dateColIndex];
  
  const dateObj = parseDate(dateStr);
  if (!dateObj) {
    continue;
  }
  
  const age = calculateAge(dateObj, referenceDate);
  
  outputItems.push({
    id: i,
    age: age,
  });
}

console.log(JSON.stringify(outputItems));