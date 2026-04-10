import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const inputPath = path.join(__dirname, 'input', 'data.json');

try {
  const rawData = fs.readFileSync(inputPath, 'utf-8');
  const data: Person[] = JSON.parse(rawData);

  const filtered = data.filter(person => person.active && person.age >= 30);
  const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(sorted));
} catch (error) {
  process.exit(1);
}