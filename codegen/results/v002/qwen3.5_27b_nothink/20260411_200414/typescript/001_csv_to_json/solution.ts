import * as fs from 'fs';
import * as path from 'path';

interface Person {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  const inputPath = path.join(__dirname, 'input', 'data.csv');
  
  // Read the CSV file
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  // Split into lines and trim to remove potential trailing newlines
  const lines = content.split('\n').filter(line => line.trim() !== '');
  
  if (lines.length < 1) {
    console.log(JSON.stringify([]));
    return;
  }
  
  // Skip header (first line)
  const dataLines = lines.slice(1);
  
  const result: Person[] = [];
  
  for (const line of dataLines) {
    // Split by comma
    const parts = line.split(',');
    
    if (parts.length < 4) {
      continue; // Skip malformed lines
    }
    
    const person: Person = {
      Name: parts[0].trim(),
      Age: parseInt(parts[1].trim(), 10),
      Email: parts[2].trim(),
      Score: parseFloat(parts[3].trim())
    };
    
    result.push(person);
  }
  
  // Output as JSON
  console.log(JSON.stringify(result));
}

main();