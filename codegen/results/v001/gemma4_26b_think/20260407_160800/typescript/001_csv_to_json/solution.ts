import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file located at 'input/data.csv'
 * and converts its rows into a JSON array of objects.
 * 
 * Expected CSV Format:
 * Name,Age,Email,Score
 * John Doe,30,john@example.com,85.5
 * 
 * Output:
 * [
 *   {
 *     "Name": "John Doe",
 *     "Age": 30,
 *     "Email": "john@example.com",
 *     "Score": 85.5
 *   }
 * ]
 */

function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    // Check if file exists to prevent throwing an unhandled error to stderr
    // which might violate the "Do not output anything else" constraint.
    if (!fs.existsSync(filePath)) {
      process.stdout.write('[]');
      return;
    }

    const fileContent = fs.readFileSync(filePath, 'utf-8');
    
    // Split into lines and remove empty lines or whitespace-only lines
    const lines = fileContent
      .split(/\r?\n/)
      .map(line => line.trim())
      .filter(line => line.length > 0);

    // If there is no data or only a header, return empty array
    if (lines.length <= 1) {
      process.stdout.write('[]');
      return;
    }

    // Skip the header row (index 0) and map the rest
    const jsonData = lines.slice(1).map(line => {
      const columns = line.split(',');

      // Mapping based on the requirements:
      // Name: string, Age: integer, Email: string, Score: float
      return {
        Name: columns[0]?.trim() || "",
        Age: parseInt(columns[1]?.trim() || "0", 10),
        Email: columns[2]?.trim() || "",
        Score: parseFloat(columns[3]?.trim() || "0")
      };
    });

    // Output the final JSON array to stdout
    process.stdout.write(JSON.stringify(jsonData, null, 2));
  } catch (error) {
    // In case of unexpected parsing errors, output an empty array to ensure valid JSON
    process.stdout.write('[]');
  }
}

main();