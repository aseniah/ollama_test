import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

try {
  const inputPath = path.join('input', 'data.json');
  const rawData = fs.readFileSync(inputPath, 'utf8');
  const data: Person[] = JSON.parse(rawData);

  const result = data
    .filter((person) => person.active && person.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(result, null, 2));
} catch (error) {
  // If there's an error (e.g., file not found), we output nothing as per instructions
  // to keep stdout clean with only the expected JSON result.
  process.exit(1);
}