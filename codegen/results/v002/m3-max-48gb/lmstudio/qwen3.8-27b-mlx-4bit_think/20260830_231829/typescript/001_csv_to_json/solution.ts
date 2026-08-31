import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

const header = lines[0].split(',');
const rows = lines.slice(1);

const result = rows.map((line) => {
  const cols = line.split(',');
  const obj: Record<string, any> = {};
  for (let i = 0; i < header.length; i++) {
    const val = cols[i];
    if (header[i] === 'Age') {
      obj['Age'] = parseInt(val, 10);
    } else if (header[i] === 'Score') {
      obj['Score'] = parseFloat(val);
    } else {
      obj[header[i]] = val;
    }
  }
  return obj;
});

console.log(JSON.stringify(result, null, 2));