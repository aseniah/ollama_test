import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.csv');

try {
  const fileContent = fs.readFileSync(filePath, 'utf-8');
  const lines = fileContent.trim().split('\n');
  
  // Skip header
  const dataLines = lines.slice(1);
  
  const result = dataLines.map(line => {
    const [name, age, email, score] = line.split(',');
    return {
      Name: name,
      Age: parseInt(age, 10),
      Email: email,
      Score: parseFloat(score)
    };
  });
  
  console.log(JSON.stringify(result));
} catch (error) {
  console.error('Error reading file:', error);
  process.exit(1);
}