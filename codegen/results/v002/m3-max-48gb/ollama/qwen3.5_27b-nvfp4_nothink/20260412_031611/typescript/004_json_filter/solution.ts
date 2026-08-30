import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

try {
  // Resolve the file path relative to the current working directory
  const filePath = path.resolve('input/data.json');
  
  // Read the file content
  const rawData = fs.readFileSync(filePath, 'utf8');
  
  // Parse JSON
  const data: Person[] = JSON.parse(rawData);

  // Filter and Sort
  const result = data
    .filter((person) => person.active === true && person.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  // Output to stdout as JSON
  console.log(JSON.stringify(result, null, 2));
} catch (error) {
  // In case of an error, output nothing or handle silently as per "Do not output anything else"
  // However, standard practice is to fail gracefully. Since the prompt implies valid input,
  // we will just exit if something goes wrong to avoid extra output.
  process.exit(1);
}