import { readFileSync } from 'fs';
import { join } from 'path';

// Define the type for the data records
interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

// Read and parse the JSON file
const filePath = join('input', 'data.json');
const rawData = readFileSync(filePath, 'utf-8');
const people: Person[] = JSON.parse(rawData);

// Filter: active must be true and age >= 30
const filtered = people.filter(p => p.active && p.age >= 30);

// Sort by name ascending
filtered.sort((a, b) => a.name.localeCompare(b.name));

// Output the result as a JSON array to stdout
console.log(JSON.stringify(filtered));