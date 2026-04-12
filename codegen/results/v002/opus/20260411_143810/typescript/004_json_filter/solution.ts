import fs from 'fs';
import path from 'path';

const filePath = path.resolve('input/data.json');
const raw = fs.readFileSync(filePath, 'utf-8');
const data: { name: string; age: number; active: boolean; score: number }[] = JSON.parse(raw);

const result = data
  .filter(r => r.active && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(result));
