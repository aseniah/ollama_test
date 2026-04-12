import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, '..', 'input', 'data.json');

fs.readFile(inputFilePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  try {
    const records = JSON.parse(data);
    const filteredRecords = records.filter(record => record.active && record.age >= 30);
    const sortedRecords = filteredRecords.sort((a, b) => a.name.localeCompare(b.name));
    console.log(JSON.stringify(sortedRecords, null, 2));
  } catch (parseErr) {
    console.error('Error parsing JSON:', parseErr);
  }
});