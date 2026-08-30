import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv',
 * parses its contents according to the specified schema,
 * and outputs a JSON array to stdout.
 */

interface UserData {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    // Check if file exists
    if (!fs.existsSync(filePath)) {
      // If the file doesn't exist, we exit silently or with error.
      // Since the requirement implies the file is provided, we assume it exists.
      process.exit(1);
    }

    // Read file content
    const content = fs.readFileSync(filePath, 'utf-8');
    
    // Split into lines and filter out empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

    if (lines.length <= 1) {
      // Only header or no data
      process.stdout.write('[]');
      return;
    }

    // Remove header
    const dataRows = lines.slice(1);

    const result: UserData[] = dataRows.map(row => {
      // Simple CSV split by comma. 
      // Note: This assumes no commas are inside quoted strings, 
      // which is fine for the provided input format.
      const columns = row.split(',');

      if (columns.length < 4) {
        throw new Error(`Invalid row format: ${row}`);
      }

      return {
        Name: columns[0].trim(),
        Age: parseInt(columns[1].trim(), 10),
        Email: columns[2].trim(),
        Score: parseFloat(columns[3].trim())
      };
    });

    // Output the JSON array to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');

  } catch (error) {
    // In a production environment, we might log to stderr.
    // For this task, we ensure valid JSON output or exit.
    process.exit(1);
  }
}

main();