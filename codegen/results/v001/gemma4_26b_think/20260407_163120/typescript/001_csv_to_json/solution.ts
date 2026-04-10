import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of each row in the output JSON.
 */
interface DataRow {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

/**
 * Main function to read the CSV file and transform it into a JSON array.
 */
function main(): void {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    // Check if the file exists before attempting to read
    if (!fs.existsSync(filePath)) {
      process.stdout.write('[]');
      return;
    }

    // Read the file content
    const csvContent = fs.readFileSync(filePath, 'utf-8');

    // Split content into lines and filter out empty lines
    const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');

    // If there's only a header or no content, return an empty array
    if (lines.length <= 1) {
      process.stdout.write('[]');
      return;
    }

    // The first line is the header, we skip it
    const dataRows = lines.slice(1);

    const jsonArray: DataRow[] = dataRows.map((line) => {
      // Split the line by comma and trim whitespace from each part
      const columns = line.split(',').map(col => col.trim());

      // Map columns to the required object structure
      // Based on the prompt requirements: Name, Age, Email, Score
      const name = columns[0] || "";
      const age = parseInt(columns[1], 10);
      const email = columns[2] || "";
      const score = parseFloat(columns[3]);

      return {
        Name: name,
        Age: isNaN(age) ? 0 : age,
        Email: email,
        Score: isNaN(score) ? 0.0 : score
      };
    });

    // Output the valid JSON array to stdout
    process.stdout.write(JSON.stringify(jsonArray));
  } catch (error) {
    // In case of a critical error, we avoid printing error logs to stdout 
    // to ensure the output remains valid JSON as per the instructions.
    // We output an empty array as a fallback.
    process.stdout.write('[]');
  }
}

// Execute the script
main();