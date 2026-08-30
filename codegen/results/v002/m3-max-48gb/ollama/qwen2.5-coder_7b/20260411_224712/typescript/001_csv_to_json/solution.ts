import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, 'input', 'data.csv');

fs.readFile(inputFilePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  const lines = data.trim().split('\n');
  const header = lines.shift()?.split(',') || [];
  const jsonArray: { [key: string]: any }[] = [];

  lines.forEach(line => {
    const values = line.split(',');
    if (values.length === header.length) {
      const obj: { [key: string]: any } = {};
      for (let i = 0; i < header.length; i++) {
        const key = header[i].trim();
        let value = values[i].trim();
        if (key === 'Age' || key === 'Score') {
          value = parseInt(value, 10);
        } else if (key === 'Email') {
          value = value.toLowerCase();
        }
        obj[key] = value;
      }
      jsonArray.push(obj);
    }
  });

  console.log(JSON.stringify(jsonArray, null, 2));
});