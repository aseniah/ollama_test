import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function solve() {
  try {
    // Define path to input file
    const inputPath = path.join(process.cwd(), 'input', 'data.json');
    
    // Read the file content
    const rawData = fs.readFileSync(inputPath, 'utf8');
    
    // Parse JSON
    const data: Record[] = JSON.parse(rawData);
    
    // Filter: active is true AND age is 30 or older
    const filtered = data.filter(item => item.active === true && item.age >= 30);
    
    // Sort: name ascending
    const sorted = filtered.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });
    
    // Output JSON array to stdout
    process.stdout.write(JSON.stringify(sorted, null, 2) + '\n');
  } catch (error) {
    // In a production environment, we would handle specific errors.
    // The requirement says "Do not output anything else", so we stay silent 
    // or let the runtime handle the exception.
    process.exit(1);
  }
}

solve();