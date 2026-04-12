import * as fs from 'fs';
import * as path from 'path';

// Define the type for the record
interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  // Resolve the path to the input file
  const filePath = path.join(__dirname, 'input', 'data.json');

  // Read the file content
  const rawData = fs.readFileSync(filePath, 'utf-8');

  // Parse the JSON data
  const data: Person[] = JSON.parse(rawData);

  // Filter: active must be true AND age must be 30 or older
  const filtered = data.filter(
    (person) => person.active === true && person.age >= 30
  );

  // Sort: by name ascending
  const sorted = filtered.sort((a, b) => {
    if (a.name < b.name) return -1;
    if (a.name > b.name) return 1;
    return 0;
  });

  // Output the result as a JSON string to stdout
  console.log(JSON.stringify(sorted));
}

main();