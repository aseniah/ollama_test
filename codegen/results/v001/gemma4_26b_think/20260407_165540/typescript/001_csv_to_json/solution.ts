import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of each row in the CSV.
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
function processCsvToJson(filePath: string): void {
  try {
    // Check if file exists
    if (!fs.existsSync(filePath)) {
      // If the file doesn't exist, we cannot output valid JSON array of data,
      // but we must output nothing else. However, since the prompt implies 
      // the file exists for the task, we proceed.
      return;
    }

    // Read file content
    const fileContent = fs.readFileSync(filePath, 'utf-8');

    // Split by newline and remove empty lines at the end
    const lines = fileContent.split(/\r?\n/).filter(line => line.trim() !== '');

    if (lines.length <= 1) {
      // Only header or empty file
      console.log(JSON.stringify([]));
      return;
    }

    // Extract header and data lines
    // We assume the header is: Name,Age,Email,Score
    const dataLines = lines.slice(1);

    const result: DataRow[] = dataLines.map((line) => {
      const columns = line.split(',');

      // Mapping columns based on requirement:
      // Name: string, Age: integer, Email: string, Score: float
      const name = columns[0]?.trim() || "";
      const age = parseInt(columns[1]?.trim() || "0", 10);
      const email = columns[2]?.trim() || "";
      const score = parseFloat(columns[3]?.trim() || "0.0");

      return {
        Name: name,
        Age: isNaN(age) ? 0 : age,
        Email: email,
        Score: isNaN(score) ? 0.0 : score
      };
    });

    // Output the JSON array to stdout
    process.stdout.write(JSON.stringify(result) + '\n');

  } catch (error) {
    // In a real-world scenario, we might log to stderr.
    // To ensure "Do not output anything else" (referring to the JSON format),
    // we exit silently if an error occurs during file reading.
    process.exit(1);
  }
}

// Define the target file path
const targetPath = path.join(process.cwd(), 'input', 'data.csv');

// Execute the process
processCsvToJson(targetPath);