import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

try {
  const filePath = path.resolve('input/data.json');
  const fileContent = fs.readFileSync(filePath, 'utf-8');
  const data: Person[] = JSON.parse(fileContent);

  const result = data
    .filter((person) => person.active && person.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(result));
} catch (error) {
  // Silent failure as per "Do not output anything else" instruction, 
  // though in a real app you might want to handle this differently.
  process.exit(1);
}