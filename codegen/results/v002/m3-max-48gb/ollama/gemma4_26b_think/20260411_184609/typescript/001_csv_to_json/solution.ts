import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv', parses its contents,
 * and outputs a JSON array to stdout.
 * 
 * The expected CSV format is:
 * Name,Age,Email,Score
 */

function solve() {
  // Define the target file path as specified in the requirements.
  // We use path.resolve to ensure we are looking for input/data.csv 
  // relative to the current working directory.
  const filePath = path.resolve('input', 'data.csv');

  try {
    // Check if the file exists before attempting to read it.
    if (!fs.existsSync(filePath)) {
      // If the file doesn't exist, we output an empty array as a fallback,
      // though the task implies the file will be present.
      process.stdout.write('[]');
      return;
    }

    // Read file content synchronously.
    const fileContent = fs.readFileSync(filePath, 'utf-8');

    // Split the content into lines and remove any trailing whitespace/empty lines.
    const lines = fileContent.trim().split(/\r?\n/);

    // If there's only a header or no lines at all, return an empty array.
    if (lines.length <= 1) {
      process.stdout.write('[]');
      return;
    }

    // The first line is the header: Name,Age,Email,Score.
    // We skip the first line and process the subsequent data rows.
    const dataRows = lines.slice(1);

    const jsonData = dataRows.map((line) => {
      // Split the line by comma to extract individual fields.
      const columns = line.split(',');

      // Ensure we have the expected number of columns.
      if (columns.length < 4) {
        return null;
      }

      // Map the columns to the specific types requested:
      // Name: string
      // Age: integer
      // Email: string
      // Score: float
      return {
        Name: columns[0].trim(),
        Age: parseInt(columns[1].trim(), 10),
        Email: columns[2].trim(),
        Score: parseFloat(columns[3].trim()),
      };
    })
    // Filter out any rows that failed to parse correctly.
    .filter((item): item is NonNullable<typeof item> => item !== null);

    // Output the final JSON array to stdout.
    // We use JSON.stringify without extra arguments to produce a single-line valid JSON.
    process.stdout.write(JSON.stringify(jsonData));
  } catch (error) {
    // In case of an unexpected error, we exit without printing anything else 
    // to maintain valid JSON output constraint, or we could print an empty array.
    // However, based on the prompt, we expect a valid file.
    process.exit(1);
  }
}

// Execute the solver.
solve();