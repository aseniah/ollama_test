import * as fs from 'fs';
import * as path from 'path';

const dataFilePath = path.join(process.cwd(), 'input', 'data.json');

const data = JSON.parse(fs.readFileSync(dataFilePath, 'utf8'));

const filtered = data
  .filter(record => record.active && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered, null, 2));