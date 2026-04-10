import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  const inputPath = path.join('input', 'data.json');
  
  // Read and parse JSON
  const rawData = fs.readFileSync(inputPath, 'utf8');
  const records: Record[] = JSON.parse(rawData);

  // Filter: active is true AND age >= 30
  const filtered = records.filter(r => r.active === true && r.age >= 30);

  // Sort by name ascending
  filtered.sort((a, b) => a.name.localeCompare(b.name));

  // Output to stdout
  console.log(JSON.stringify(filtered));
}

main();