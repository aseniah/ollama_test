import * as fs from 'fs';
import * as path from 'path';

interface UserRecord {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function solve() {
  const filePath = path.join('input', 'data.csv');

  try {
    // Read the file content
    const content = fs.readFileSync(filePath, 'utf-8');
    
    // Split by lines and filter out empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    // Remove the header row
    const rows = lines.slice(1);
    
    const result: UserRecord[] = rows.map(row => {
      const columns = row.split(',');
      
      return {
        Name: columns[0],
        Age: parseInt(columns[1], 10),
        Email: columns[2],
        Score: parseFloat(columns[3])
      };
    });

    // Output the final JSON array to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
  } catch (error) {
    // In a production environment, we would handle this more gracefully.
    // For this task, we assume the file exists as described.
    process.exit(1);
  }
}

solve();