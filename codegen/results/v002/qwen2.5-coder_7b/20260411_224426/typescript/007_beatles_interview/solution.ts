import * as fs from 'fs';
import { join } from 'path';

const inputPath = join(__dirname, 'input', 'input.csv');
const expectedFormatPath = join(__dirname, 'input', 'expected_format.json');

// Read the CSV file
const csvData = fs.readFileSync(inputPath, 'utf8');

// Parse the CSV data
const lines = csvData.split('\n').filter(line => line.trim() !== '');
const headers = lines[0].split(',').map(header => header.trim());
const rows = lines.slice(1).map(line => {
  const values = line.split(',').map(value => value.trim().replace(/"/g, ''));
  return headers.reduce((acc, header, index) => {
    acc[header] = values[index];
    return acc;
  }, {} as any);
});

// Calculate ages and prepare the output
const output: any[] = rows.map(row => {
  const birthday = new Date(row.Birthday).toISOString().split('T')[0];
  const age = calculateAge(birthday, '2025-07-01');
  
  const relatives: any[] = [];
  if (row.Father) relatives.push({ FirstName: row.Father, LastName: '', Relationship: 'Father' });
  if (row.Mother) relatives.push({ FirstName: row.Mother, LastName: '', Relationship: 'Mother' });
  if (row.Brother) relatives.push({ FirstName: row.Brother, LastName: '', Relationship: 'Brother' });
  if (row.Sister) relatives.push({ FirstName: row.Sister, LastName: '', Relationship: 'Sister' });

  return {
    FirstName: row.Name.split(' ')[0],
    LastName: row.Name.split(' ').slice(1).join(' '),
    Birthday,
    Age: age,
    Relatives
  };
});

// Write the output to stdout
process.stdout.write(JSON.stringify(output, null, 2));

// Function to calculate age
function calculateAge(birthday: string, currentDate: string): number {
  const birthDate = new Date(birthday);
  const current = new Date(currentDate);
  
  let age = current.getFullYear() - birthDate.getFullYear();
  const monthDiff = current.getMonth() - birthDate.getMonth();

  if (monthDiff < 0 || (monthDiff === 0 && current.getDate() < birthDate.getDate())) {
    age--;
  }

  return age;
}