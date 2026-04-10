import * as fs from 'fs';
import * as path from 'path';

interface DataRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = process.argv[2] || 'input/data.json';
const dataPath = path.join(process.cwd(), filePath);

const fileContent = fs.readFileSync(dataPath, 'utf-8');
const records: DataRecord[] = JSON.parse(fileContent);

const filtered = records.filter(record => record.active && record.age >= 30);

filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered));