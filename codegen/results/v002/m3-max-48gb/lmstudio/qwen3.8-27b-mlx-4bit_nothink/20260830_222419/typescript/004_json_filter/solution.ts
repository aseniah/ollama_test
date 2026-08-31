import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const inputPath = path.join('input', 'data.json');
const data: Record[] = JSON.parse(fs.readFileSync(inputPath, 'utf-8'));

const filtered = data
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(filtered, null, 2));