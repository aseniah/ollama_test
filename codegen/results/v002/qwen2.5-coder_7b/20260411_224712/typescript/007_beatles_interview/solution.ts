import { readFileSync } from 'fs';
import path from 'path';

// Define the input and output file paths
const inputFilePath = path.join(__dirname, 'input', 'input.csv');
const expectedFormatFilePath = path.join(__dirname, 'input', 'expected_format.json');

// Read the input CSV file
const inputFileContent = readFileSync(inputFilePath, 'utf8');
const lines = inputFileContent.trim().split('\n');
const headers = lines[0].split(',');

// Function to calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  const birthDate = new Date(birthday);
  const targetDate = new Date(2025, 6, 1); // July 1, 2025
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = targetDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && targetDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

// Process each line to transform it into the desired output format
const result: any[] = lines.slice(1).map(line => {
  const values = line.split(',');
  const person = {
    FirstName: values[0].split(' ').slice(0, -1).join(' '),
    LastName: values[0].split(' ').pop(),
    Birthday: new Date(values[1]).toISOString().split('T')[0],
    Age: calculateAge(values[1]),
    Relatives: []
  };

  if (values[4]) person.Relatives.push({ FirstName: values[4], LastName: '', Relationship: 'Father' });
  if (values[5]) person.Relatives.push({ FirstName: values[5], LastName: '', Relationship: 'Mother' });
  if (values[6]) person.Relatives.push({ FirstName: values[6], LastName: '', Relationship: 'Brother' });
  if (values[7]) person.Relatives.push({ FirstName: values[7], LastName: '', Relationship: 'Sister' });

  return person;
});

// Write the result to stdout
process.stdout.write(JSON.stringify(result, null, 2));