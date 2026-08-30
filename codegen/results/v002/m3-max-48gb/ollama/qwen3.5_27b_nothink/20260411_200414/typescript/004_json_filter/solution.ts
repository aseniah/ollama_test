import * as fs from 'fs';
import * as path from 'path';

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

// Define the input file path relative to the current working directory
const inputFile = path.join('input', 'data.json');

try {
  // Read the file synchronously
  const rawData = fs.readFileSync(inputFile, 'utf-8');

  // Parse the JSON data
  const data: Person[] = JSON.parse(rawData);

  // Filter records where active is true and age >= 30
  const filtered = data.filter(
    (person) => person.active === true && person.age >= 30
  );

  // Sort the filtered records by name in ascending order
  filtered.sort((a, b) => a.name.localeCompare(b.name));

  // Output the result as a JSON string to stdout
  console.log(JSON.stringify(filtered));
} catch (error) {
  // Handle potential errors (e.g., file not found, invalid JSON)
  console.error('Error processing file:', error);
  process.exit(1);
}