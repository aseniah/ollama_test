import * as fs from 'fs';
import * as path from 'path';

const inputFile = process.argv[2];
const inputPath = path.join(__dirname, 'input', inputFile);

if (!inputFile) {
  console.error('Please provide a file name as an argument.');
  process.exit(1);
}

fs.readFile(inputPath, 'utf8', (err, data) => {
  if (err) {
    console.error(`Error reading the file: ${err.message}`);
    process.exit(1);
  }

  const rows = data.split('\n').filter(row => row.trim() !== '');
  const header = rows[0].split(',');
  const result: { [key: string]: any }[] = [];

  for (let i = 1; i < rows.length; i++) {
    const row = rows[i].split(',');
    const obj: { [key: string]: any } = {};

    for (let j = 0; j < header.length; j++) {
      const key = header[j];
      const value = row[j];

      if (key === 'Age') {
        obj[key] = parseInt(value, 10);
      } else if (key === 'Score') {
        obj[key] = parseFloat(value);
      } else {
        obj[key] = value;
      }
    }

    result.push(obj);
  }

  console.log(JSON.stringify(result, null, 2));
});