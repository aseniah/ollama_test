import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

const header = lines[0].split(',');
const result = lines.slice(1).map(line => {
  const cols = line.split(',');
  const obj: Record<string, any> = {};
  for (let i = 0; i < header.length; i++) {
    const key = header[i].trim();
    const value = cols[i].trim();
    if (key === 'Age') {
      obj[key] = parseInt(value, 10);
    } else if (key === 'Score') {
      obj[key] = parseFloat(value);
    } else {
      obj[key] = value;
    }
  }
  return obj;
});

process.stdout.write(JSON.stringify(result, null, 2) + '\n');