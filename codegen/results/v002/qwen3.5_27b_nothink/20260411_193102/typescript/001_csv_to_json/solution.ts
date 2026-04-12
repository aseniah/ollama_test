import * as fs from 'fs';
import * as path from 'path';

interface DataEntry {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');
  
  try {
    const content = fs.readFileSync(filePath, 'utf-8');
    const lines = content.trim().split('\n');

    if (lines.length < 2) {
      console.log('[]');
      return;
    }

    // Skip header (first line)
    const dataLines = lines.slice(1);
    
    const result: DataEntry[] = dataLines.map(line => {
      // Handle CSV parsing that might have commas within quoted fields
      // For this specific input format, simple split is sufficient as there are no commas in fields
      const parts = line.split(',');
      
      if (parts.length !== 4) {
        throw new Error(`Invalid CSV line format: ${line}`);
      }

      return {
        Name: parts[0],
        Age: parseInt(parts[1], 10),
        Email: parts[2],
        Score: parseFloat(parts[3])
      };
    });

    // Output as valid JSON to stdout
    console.log(JSON.stringify(result, null, 2));
  } catch (error) {
    console.error('Error:', (error as Error).message);
    process.exit(1);
  }
}

main();