import * as fs from 'fs';
import * as path from 'path';

const inputFile = path.join(process.cwd(), 'input', 'data.json');

let rawData: string;
try {
  rawData = fs.readFileSync(inputFile, 'utf-8');
} catch (error) {
  console.error(`Failed to read file: ${inputFile}`);
  process.exit(1);
}

let data: Array<{ name: string; age: number; active: boolean; score: number }> = [];
try {
  data = JSON.parse(rawData);
} catch (error) {
  console.error('Failed to parse JSON');
  process.exit(1);
}

const filtered = data.filter(record => record.active === true && record.age >= 30);
const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(sorted));