import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(process.cwd(), 'input', 'data.csv');

fs.readFile(inputFilePath, 'utf8', (err, data) => {
  if (err) {
    console.error(`Error reading file: ${err.message}`);
    return;
  }

  const lines = data.split('\n');
  const header = lines[0].split(',');
  const result: any[] = [];

  for (let i = 1; i < lines.length; i++) {
    if (lines[i]) {
      const row = lines[i].split(',');
      const obj: any = {};
      for (let j = 0; j < header.length; j++) {
        switch (header[j]) {
          case 'Name':
            obj.Name = row[j];
            break;
          case 'Age':
            obj.Age = parseInt(row[j], 10);
            break;
          case 'Email':
            obj.Email = row[j];
            break;
          case 'Score':
            obj.Score = parseFloat(row[j]);
            break;
        }
      }
      result.push(obj);
    }
  }

  console.log(JSON.stringify(result, null, 2));
});