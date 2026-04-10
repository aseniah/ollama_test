import * as fs from 'fs';

// Read input file
const inputContent = fs.readFileSync('input/data.json', 'utf-8');
let data: any;
try {
  data = JSON.parse(inputContent);
} catch (e) {
  console.error('Failed to parse input JSON');
  process.exit(1);
}

if (!Array.isArray(data)) {
  console.error('Input data must be an array');
  process.exit(1);
}

const result = data.filter((record: any) => {
  return (
    typeof record.active === 'boolean' && record.active === true &&
    typeof record.age === 'number' && record.age >= 30
  );
});

// Sort by name ascending
result.sort((a: any, b: any) => {
  if (a.name === undefined || b.name === undefined) {
    return 0;
  }
  return a.name.localeCompare(b.name);
});

console.log(JSON.stringify(result, null, 2));