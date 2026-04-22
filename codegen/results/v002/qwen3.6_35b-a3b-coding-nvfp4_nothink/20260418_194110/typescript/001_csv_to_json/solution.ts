import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');
const header = lines[0].split(',');

const result: Array<Record<string, string | number>> = [];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  if (!line) continue;
  const values = line.split(',');
  const obj: Record<string, string | number> = {};
  obj['Name'] = values[0];
  obj['Age'] = parseInt(values[1], 10);
  obj['Email'] = values[2];
  obj['Score'] = parseFloat(values[3]);
  result.push(obj);
}

console.log(JSON.stringify(result));