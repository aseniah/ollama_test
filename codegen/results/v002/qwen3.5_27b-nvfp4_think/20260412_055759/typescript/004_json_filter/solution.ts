import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const dataPath = path.join('input', 'data.json');
const rawData = fs.readFileSync(dataPath, 'utf-8');
const records: Record[] = JSON.parse(rawData);

const filteredRecords = records.filter(r => r.active && r.age >= 30);

filteredRecords.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filteredRecords, null, 2));