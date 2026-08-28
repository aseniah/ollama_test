import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.json');

try {
  const data = fs.readFileSync(filePath, 'utf8');
  const records = JSON.parse(data);
  
  const filteredRecords = records
    .filter((record: any) => record.active === true && record.age >= 30)
    .sort((a: any, b: any) => a.name.localeCompare(b.name));
  
  console.log(JSON.stringify(filteredRecords));
} catch (error) {
  console.error('Error reading or processing the file:', error);
}