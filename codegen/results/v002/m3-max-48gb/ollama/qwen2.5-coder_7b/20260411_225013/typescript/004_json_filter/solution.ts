import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = join(__dirname, 'input', 'data.json');

try {
  const fileContent = readFileSync(filePath, 'utf-8');
  const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(fileContent);

  const filteredData = data
    .filter(item => item.active && item.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(filteredData));
} catch (error) {
  console.error('Error reading or processing the file:', error);
}