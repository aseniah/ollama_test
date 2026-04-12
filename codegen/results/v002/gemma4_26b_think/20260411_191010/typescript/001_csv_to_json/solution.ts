import * as fs from 'fs';
import * as path from 'path';

/**
 * This program reads a CSV file located at 'input/data.csv' and converts
 * its content into a JSON array of objects.
 * 
 * The CSV structure is expected to be:
 * Name,Age,Email,Score
 * 
 * Output is written to stdout as a valid JSON array.
 */

function solve() {
  const filePath = path.join('input', 'data.csv');

  try {
    // Check if the input file exists to prevent errors.
    if (!fs.existsSync(filePath)) {
      process.stdout.write('[]');
      return;
    }

    // Read file content synchronously.
    const fileContent = fs.readFileSync(filePath, 'utf-8');

    // Split content by newline characters (handling both \n and \r\n).
    const lines = fileContent.split(/\r?\n/);

    // Filter out any empty lines that might occur at the end of the file.
    const validLines = lines.filter(line => line.trim() !== '');

    if (validLines.length <= 1) {
      // If only the header exists or the file is empty, return an empty array.
      process.stdout.write('[]');
      return;
    }

    // Remove the header row (the first line).
    const dataRows = validLines.slice(1);

    const result: Array<{
      Name: string;
      Age: number;
      Email: string;
      Score: number;
    }> = dataRows.map(row => {
      const columns = row.split(',');

      // The problem specifies:
      // Name (string), Age (integer), Email (string), Score (float)
      const name = columns[0]?.trim() || '';
      const age = parseInt(columns[1]?.trim() || '0', 10);
      const email = columns[2]?.trim() || '';
      const score = parseFloat(columns[3]?.trim() || '0.0');

      return {
        Name: name,
        Age: age,
        Email: email,
        Score: score
      };
    });

    // Output the JSON array to stdout.
    // Using JSON.stringify without indentation to ensure strict adherence 
    // to "do not output anything else".
    process.stdout.write(JSON.stringify(result));

  } catch (error) {
    // In case of unexpected errors (e.g., permission issues),
    // we output an empty array to satisfy the "valid JSON" requirement.
    process.stdout.write('[]');
  }
}

// Execute the solution.
solve();