import * as fs from 'fs';

const inputPath = 'input/data.csv';
const content = fs.readFileSync(inputPath, 'utf-8');
const lines = content.trim().split('\n');
const header = lines[0].split(',');
const data: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  const obj: Record<string, any> = {};
  for (let j = 0; j < header.length; j++) {
    const key = header[j];
    const val = values[j];
    if (key === 'Age') {
      obj['Age'] = parseInt(val, 10);
    } else if (key === 'Score') {
      obj['Score'] = parseFloat(val);
    } else {
      obj[key] = val;
    }
  }
  data.push(obj);
}

console.log(JSON.stringify(data));