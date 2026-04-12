import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv', parses its content,
 * and outputs a JSON array to stdout.
 */

interface UserData {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function solve() {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    // Check if file exists
    if (!fs.existsSync(filePath)) {
      // If the specific path doesn't exist, we can't proceed.
      // In a real environment, we'd throw an error, but here we exit silently 
      // to satisfy "Do not output anything else" if the input is missing.
      process.exit(1);
    }

    const content = fs.readFileSync(filePath, 'utf-8');
    
    // Split content into lines and remove empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

    if (lines.length < 2) {
      // Only header or no content
      console.log(JSON.stringify([]));
      return;
    }

    // Extract header to identify columns (though task specifies fixed mapping)
    const headers = lines[0].split(',').map(h => h.trim());
    
    const result: UserData[] = [];

    // Iterate through data rows (skipping header)
    for (let i = 1; i < lines.length; i++) {
      const columns = lines[i].split(',').map(c => c.trim());
      
      // Ensure we have the expected number of columns
      if (columns.length === 4) {
        const row: UserData = {
          Name: columns[0],
          Age: parseInt(columns[1], 10),
          Email: columns[2],
          Score: parseFloat(columns[3])
        };
        result.push(row);
      }
    }

    // Output the final JSON array to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');

  } catch (error) {
    // If any error occurs, we exit. 
    // The requirement is "Do not output anything else" besides valid JSON.
    process.exit(1);
  }
}

solve();