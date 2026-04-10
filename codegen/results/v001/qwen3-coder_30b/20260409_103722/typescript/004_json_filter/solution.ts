import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.json');

try {
  const data = fs.readFileSync(filePath, 'utf-8');
  const records = JSON.parse(data);
  
  const filtered = records
    .filter((record: any) => record.active && record.age >= 30)
    .sort((a: any, b: any) => a.name.localeCompare(b.name));
  
  console.log(JSON.stringify(filtered));
} catch (error) {
  // Handle file reading or parsing errors silently as per requirements
}