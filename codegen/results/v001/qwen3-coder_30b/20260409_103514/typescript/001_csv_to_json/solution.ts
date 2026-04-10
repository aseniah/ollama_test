import fs from 'fs';
import path from 'path';

const filePath = path.join('input', 'data.csv');

try {
  const fileContent = fs.readFileSync(filePath, 'utf-8');
  const lines = fileContent.trim().split('\n');
  
  // Skip header
  const dataLines = lines.slice(1);
  
  const result = dataLines.map(line => {
    const [name, age, email, score] = line.split(',');
    return {
      Name: name.trim(),
      Age: parseInt(age.trim(), 10),
      Email: email.trim(),
      Score: parseFloat(score.trim())
    };
  });
  
  console.log(JSON.stringify(result));
} catch (error) {
  console.error('Error reading file:', error);
  process.exit(1);
}