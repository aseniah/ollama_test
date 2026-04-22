import fs from 'fs';
import path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join(process.cwd(), 'input', 'data.json');
const records: Record[] = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

const result = records
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));