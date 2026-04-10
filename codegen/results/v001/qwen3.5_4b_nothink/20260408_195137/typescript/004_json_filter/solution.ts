import fs from 'fs';

// Read input file
const inputPath = 'input/data.json';
let data: any[];
try {
  const content = fs.readFileSync(inputPath, 'utf8');
  data = JSON.parse(content);
} catch (error) {
  console.error('Error reading or parsing input file:', error);
  process.exit(1);
}

// Filter records: active is true AND age >= 30
const filteredData = data.filter(
  (record: any) =>
    record.active === true &&
    (typeof record.age === 'number' && record.age >= 30)
);

// Sort by name ascending
filteredData.sort((a: any, b: any) => {
  return a.name.localeCompare(b.name);
});

// Output to stdout
console.log(JSON.stringify(filteredData));