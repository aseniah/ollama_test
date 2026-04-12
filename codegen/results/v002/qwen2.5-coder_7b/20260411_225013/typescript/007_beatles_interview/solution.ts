import fs from 'fs';
import path from 'path';

const inputCsvPath = path.join(__dirname, 'input', 'input.csv');
const expectedJsonPath = path.join(__dirname, 'input', 'expected_format.json');

// Read the CSV file
const csvContent = fs.readFileSync(inputCsvPath, 'utf-8');
const lines = csvContent.split('\n').slice(1); // Skip header

// Parse the CSV data
const data: {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Died?: string;
  Father?: string;
  Mother?: string;
  Brother?: string;
  Sister?: string;
}[] = lines.map(line => {
  const parts = line.split(',');
  return {
    FirstName: parts[0].split(' ')[0],
    LastName: parts[0].split(' ')[1] || '',
    Birthday: parts[1],
    Died: parts[2],
    Father: parts[3],
    Mother: parts[4],
    Brother: parts[5],
    Sister: parts[6]
  };
});

// Calculate ages
const targetDate = new Date('2025-07-01');
data.forEach(person => {
  const birthDate = new Date(person.Birthday);
  const age = targetDate.getFullYear() - birthDate.getFullYear();
  if (targetDate.getMonth() < birthDate.getMonth() || 
      (targetDate.getMonth() === birthDate.getMonth() && targetDate.getDate() < birthDate.getDate())) {
    person.Age = age - 1;
  } else {
    person.Age = age;
  }
});

// Create relatives array
const relatives: any[] = [];
data.forEach(person => {
  const relativesForPerson: any[] = [];
  if (person.Father) {
    relativesForPerson.push({ FirstName: person.Father, LastName: '', Relationship: 'Father' });
  }
  if (person.Mother) {
    relativesForPerson.push({ FirstName: person.Mother, LastName: '', Relationship: 'Mother' });
  }
  if (person.Brother) {
    relativesForPerson.push({ FirstName: person.Brother, LastName: '', Relationship: 'Brother' });
  }
  if (person.Sister) {
    relativesForPerson.push({ FirstName: person.Sister, LastName: '', Relationship: 'Sister' });
  }
  relatives.push(...relativesForPerson);
});

// Map the data to the expected format
const result = data.map(person => ({
  ...person,
  Relatives: relatives.filter(relative => 
    relative.FirstName === person.Father ||
    relative.FirstName === person.Mother ||
    relative.FirstName === person.Brother ||
    relative.FirstName === person.Sister)
}));

// Remove undefined properties and sort by FirstName and LastName
const sortedResult = result.map(person => ({
  ...person,
  Relatives: person.Relatives.filter(relative => relative.FirstName !== undefined && relative.LastName !== undefined),
  Age: Number(person.Age),
  Birthday: new Date(person.Birthday).toISOString().split('T')[0]
})).sort((a, b) => {
  if (a.FirstName === b.FirstName) return a.LastName.localeCompare(b.LastName);
  return a.FirstName.localeCompare(b.FirstName);
});

// Write the result to stdout
process.stdout.write(JSON.stringify(sortedResult, null, 2));