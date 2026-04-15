import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of each record in the CSV.
 */
interface Record {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  try {
    // Define the path to the input file as specified in the prompt
    const filePath = path.join('input', 'data.csv');

    // Read the file content synchronously using utf8 encoding
    const content = fs.readFileSync(filePath, 'utf8');

    // Split content into lines and filter out any trailing empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

    // Ensure there is at least a header and one row of data
    if (lines.length < 1) {
      process.stdout.write(JSON.stringify([]));
      return;
    }

    // Extract the data rows, skipping the first line (header)
    const dataRows = lines.slice(1);

    const results: Record[] = dataRows.map(line => {
      // Split by comma. Note: This assumes simple CSV without quoted delimiters.
      const [name, ageStr, email, scoreStr] = line.split(',');

      return {
        Name: name,
        Age: parseInt(ageStr, 10),
        Email: email,
        Score: parseFloat(scoreStr),
      };
    });

    // Output the final array as a JSON string to stdout
    process.stdout.write(JSON.stringify(results, null, 2));
  } catch (error) {
    // The prompt specifies "Do not output anything else", but in a real 
    // production environment, we would handle errors here.
    // Given the strict constraints, we fail silently or let the process crash.
    process.exit(1);
  }
}

// Execute the program
main();