import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join('input', 'data.json');

try {
  const rawData = fs.readFileSync(filePath, 'utf-8');
  const records: Record[] = JSON.parse(rawData);

  const filteredRecords = records
    .filter((record) => record.active === true && record.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(filteredRecords));
} catch (error) {
  // In case of error, we should not output anything else as per instructions.
  // However, if the file is missing or invalid, the process will likely exit with an error.
  process.exit(1);
}