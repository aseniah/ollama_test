import { readFileSync } from 'fs';

const content = readFileSync('./input/data.csv', 'utf-8');

const lines = content.split('\n').filter(line => line.trim());

if (lines.length === 0) {
  process.stdout.write('[]\n');
  process.exit(0);
}

const headerLine = lines[0].split(',');
const columnNameToIndex: Map<string, number> = new Map();

for (let i = 0; i < headerLine.length; i++) {
  const colName = headerLine[i].trim().toLowerCase();
  if (colName === 'name') columnNameToIndex.set('name', i);
  else if (colName === 'age') columnNameToIndex.set('age', i);
  else if (colName === 'email') columnNameToIndex.set('email', i);
  else if (colName === 'score') columnNameToIndex.set('score', i);
}

const data: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const columns = lines[i].split(',');
  
  let entry: any = {};
  for (const [name, index] of columnNameToIndex.entries()) {
    if (index >= 0 && index < columns.length) {
      value = columns[index];
      if (name === 'age') entry['Age'] = parseInt(value.trim(), 10);
      else if (name === 'score') entry['Score'] = parseFloat(value.trim());
      else if (name === 'name' || name === 'email') entry[name] = value.trim();
    }
  }
  
  data.push(entry as any);
}

console.log(JSON.stringify(data));