import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

const header = lines[0].split(',');
const rows: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  if (line === '') continue;
  
  const values = line.split(',');
  const row: any = {};
  
  for (let j = 0; j < header.length; j++) {
    const field = header[j].trim();
    const value = values[j].trim();
    
    switch (field) {
      case 'Name':
        row.Name = value;
        break;
      case 'Age':
        row.Age = parseInt(value, 10);
        break;
      case 'Email':
        row.Email = value;
        break;
      case 'Score':
        row.Score = parseFloat(value);
        break;
      default:
        row[field] = value;
    }
  }
  
  rows.push(row);
}

process.stdout.write(JSON.stringify(rows, null, 2) + '\n');