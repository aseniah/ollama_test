import * as fs from 'fs';
import { parse } from 'csv-parse/sync';

const inputPath = process.argv[2] || 'input/data.csv';
const outputPath = process.argv[3];

const data = fs.readFileSync(inputPath, 'utf-8');
const records = parse(data, {
  columns: true,
  skip_empty_lines: true,
});

const result = records.map((row) => ({
  name: String(row.name),
  age: parseInt(row.age, 10),
  email: String(row.email),
  score: parseFloat(row.score),
}));

console.log(JSON.stringify(result, null, 2));

if (outputPath) {
  fs.writeFileSync(outputPath, JSON.stringify(result, null, 2));
}