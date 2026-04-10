import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = join('input', 'data.json');

try {
  const fileContent = readFileSync(filePath, 'utf-8');
  const data = JSON.parse(fileContent);
  
  const filteredData = data
    .filter((item: any) => item.active === true && item.age >= 30)
    .sort((a: any, b: any) => a.name.localeCompare(b.name));
  
  console.log(JSON.stringify(filteredData));
} catch (error) {
  // Do nothing as per requirements
}