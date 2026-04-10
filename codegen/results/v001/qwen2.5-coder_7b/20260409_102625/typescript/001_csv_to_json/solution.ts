import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, 'input', 'data.csv');

fs.readFile(inputFilePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  const lines = data.split('\n');
  const header = lines[0].split(',');
  const result: { [key: string]: any }[] = [];

  for (let i = 1; i < lines.length; i++) {
    if (lines[i].trim() === '') continue;
    const row = lines[i].split(',');
    const obj: { [key: string]: any } = {};

    for (let j = 0; j < header.length; j++) {
      const key = header[j].trim();
      let value = row[j].trim();

      if (key === 'Age') {
        value = parseInt(value, 10);
      } else if (key === 'Score') {
        value = parseFloat(value);
      }

      obj[key] = value;
    }

    result.push(obj);
  }

  console.log(JSON.stringify(result, null, 2));
});