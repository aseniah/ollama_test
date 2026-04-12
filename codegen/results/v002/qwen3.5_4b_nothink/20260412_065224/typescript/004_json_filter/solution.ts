import * as fs from 'fs';

// Read input file
const inputFile = 'input/data.json';
let data: any[];
try {
  const fileContent = fs.readFileSync(inputFile, 'utf-8');
  data = JSON.parse(fileContent);
} catch (error) {
  console.error(`Error reading ${inputFile}: ${error}`);
  process.exit(1);
}

// Filter and transform
const filtered = data.filter(record => {
  return (
    record.active === true &&
    record.age >= 30
  );
});

// Sort by name ascending
const sorted = filtered.sort((a, b) => {
  return a.name.localeCompare(b.name);
});

// Output to stdout
console.log(JSON.stringify(sorted, null, 2));