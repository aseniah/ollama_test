import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv' and converts it to a JSON array.
 * It assumes the CSV has a header row and the columns are Name, Age, Email, Score.
 */

function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    // Check if file exists
    if (!fs.existsSync(filePath)) {
      process.exit(1);
    }

    const content = fs.readFileSync(filePath, 'utf-8');
    
    // Split by lines and filter out empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

    if (lines.length <= 1) {
      // Only header or empty file
      console.log(JSON.stringify([]));
      return;
    }

    // Remove header
    const dataRows = lines.slice(1);

    const result = dataRows.map((line) => {
      // Split by comma, handling potential basic CSV structure
      // Note: This does not handle complex quoted CSVs with commas inside quotes,
      // as per the requirement for a standard simple CSV parsing.
      const columns = line.split(',').map(col => col.trim());

      if (columns.length < 4) {
        // Skip malformed rows or handle as needed
        return null;
      }

      const name = columns[0];
      const age = parseInt(columns[1], 10);
      const email = columns[2];
      const score = parseFloat(columns[3]);

      return {
        Name: name,
        Age: isNaN(age) ? 0 : age,
        Email: email,
        Score: isNaN(score) ? 0.0 : score
      };
    }).filter((item): item is NonNullable<typeof item> => item !== null);

    // Output valid JSON to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');

  } catch (error) {
    // On error, we exit silently or with error code to ensure only valid JSON goes to stdout
    process.exit(1);
  }
}

main();