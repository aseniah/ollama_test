import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.join(process.cwd(), 'input', 'data.json');
const filteredRecords: any[] = [];

try {
  const rawData = fs.readFileSync(inputFile, 'utf8');
  const records = JSON.parse(rawData);

  // Filter records where active is true and age >= 30
  const filteredRecords = records.filter(record => 
    record.active === true && 
    (typeof record.age !== 'undefined' && Number.isInteger(record.age) || typeof record.age === 'number') &&
    record.age >= 30
  );

  // Sort by name ascending
  filteredRecords.sort((a, b) => a.name.localeCompare(b.name));

  // Output to stdout as JSON array
  console.log(JSON.stringify(filteredRecords, null, 2));
} catch (error) {
  console.error('Error reading or parsing file:', error.message);
  process.exit(1);
}