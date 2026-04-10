import * as fs from 'fs';
import * as path from 'path';

/**
 * Problem: Transform CSV data into a specific JSON format.
 * Constraints: 
 * - Calculate age as of July 1, 2025.
 * - Output only the JSON array to stdout.
 * - No external npm packages.
 */

interface InputRow {
  [key: string]: string;
}

interface OutputRecord {
  id: string;
  fullName: string;
  age: number;
  registrationDate: string;
  status: 'active' | 'inactive';
}

function parseCSV(content: string): InputRow[] {
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  if (lines.length === 0) return [];

  const headers = lines[0].split(',').map(h => h.trim());
  const rows: InputRow[] = [];

  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',').map(v => v.trim());
    const row: InputRow = {};
    headers.forEach((header, index) => {
      row[header] = values[index] || '';
    });
    rows.push(row);
  }

  return rows;
}

function calculateAge(birthDateStr: string, referenceDate: Date): number {
  const birthDate = new Date(birthDateStr);
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'input.csv');
  
  if (!fs.existsSync(inputPath)) {
    process.exit(0);
  }

  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  const rows = parseCSV(csvContent);

  const referenceDate = new Date('2025-07-01T00:00:00');

  /**
   * Transformation logic inferred from the task:
   * We assume the CSV contains columns like 'id', 'firstName', 'lastName', 'birthDate', 'regDate', 'isActive'
   * and the JSON needs 'fullName', 'age', 'registrationDate', 'status'.
   */
  const result: OutputRecord[] = rows.map(row => {
    const firstName = row['firstName'] || '';
    const lastName = row['lastName'] || '';
    const birthDateStr = row['birthDate'] || '';
    const regDateStr = row['regDate'] || '';
    const isActive = row['isActive']?.toLowerCase() === 'true';

    return {
      id: row['id'] || '',
      fullName: `${firstName} ${lastName}`.trim(),
      age: birthDateStr ? calculateAge(birthDateStr, referenceDate) : 0,
      registrationDate: regDateStr,
      status: isActive ? 'active' : 'inactive'
    };
  });

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();