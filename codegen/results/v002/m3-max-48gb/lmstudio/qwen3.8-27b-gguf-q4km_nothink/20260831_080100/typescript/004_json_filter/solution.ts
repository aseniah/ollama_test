import * as fs from 'fs';
import * as path from 'path';

// Read the input file
const inputPath = path.join('input', 'data.json');
const data = JSON.parse(fs.readFileSync(inputPath, 'utf8'));

// Filter and sort
const result = data
  .filter((record: any) => record.active === true && record.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

// Output
console.log(JSON.stringify(result, null, 2));