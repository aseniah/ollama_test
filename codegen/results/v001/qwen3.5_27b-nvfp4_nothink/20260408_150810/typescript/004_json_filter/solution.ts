import * as fs from 'fs';
import * as path from 'path';

// Define the interface for the data objects
interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

// Resolve the file path
const filePath = path.join('input', 'data.json');

try {
  // Read and parse the JSON file
  const rawData = fs.readFileSync(filePath, 'utf8');
  const data: Person[] = JSON.parse(rawData);

  // Filter records where active is true and age is >= 30
  const filtered = data.filter(person => person.active && person.age >= 30);

  // Sort the filtered records by name ascending
  filtered.sort((a, b) => a.name.localeCompare(b.name));

  // Output the result as JSON to stdout
  console.log(JSON.stringify(filtered));
} catch (error) {
  // If there's an error (e.g., file not found), exit with an error code
  process.exit(1);
}