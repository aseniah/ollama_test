import * as fs from 'fs';
import * as path from 'path';

interface UserRecord {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function solve() {
  try {
    // Path to the input file
    const filePath = path.join('input', 'data.csv');
    
    // Read the file content as a string
    const content = fs.readFileSync(filePath, 'utf8');
    
    // Split by newline and filter out any empty lines (e.g., at the end of the file)
    const lines = content.split(/\r?\n/).filter(line => line.trim().length > 0);
    
    if (lines.length === 0) {
      console.log(JSON.stringify([]));
      return;
    }

    // Remove the header row
    const dataLines = lines.slice(1);
    
    const results: UserRecord[] = dataLines.map(line => {
      const columns = line.split(',');
      
      return {
        Name: columns[0]?.trim() || '',
        Age: parseInt(columns[1]?.trim() || '0', 10),
        Email: columns[2]?.trim() || '',
        Score: parseFloat(columns[3]?.trim() || '0')
      };
    });

    // Output the final array as JSON
    process.stdout.write(JSON.stringify(results, null, 2) + '\n');
  } catch (error) {
    // In a real scenario, we might handle errors differently, 
    // but here we'll just exit if the file cannot be read.
    process.exit(1);
  }
}

solve();