import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

try {
  const filePath = path.join(process.cwd(), 'input', 'data.json');
  const fileContent = fs.readFileSync(filePath, 'utf-8');
  const data: Person[] = JSON.parse(fileContent);

  const filteredAndSorted = data
    .filter((person) => person.active && person.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(filteredAndSorted));
} catch (error) {
  process.exit(1);
}