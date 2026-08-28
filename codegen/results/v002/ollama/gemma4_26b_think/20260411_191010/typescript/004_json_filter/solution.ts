import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of each record in the input JSON.
 */
interface DataRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * The main function to read the input file, filter the records,
 * sort them by name, and output the result to stdout.
 */
function processData() {
  // The path to the input file as specified in the requirements.
  const inputPath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if the file exists before attempting to read.
    if (!fs.existsSync(inputPath)) {
      return;
    }

    // Read the file content synchronously.
    const fileContent = fs.readFileSync(inputPath, 'utf8');

    // Parse the JSON string into an array of DataRecord objects.
    const records: DataRecord[] = JSON.parse(fileContent);

    // 1. Filter: Only records where 'active' is true and 'age' is 30 or older.
    // 2. Sort: Sort the remaining records by 'name' in ascending order.
    const processedRecords = records
      .filter((record) => record.active === true && record.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the resulting array to stdout as a formatted JSON string.
    // We use null, 2 to provide a readable, pretty-printed JSON output.
    // The newline at the end is standard for terminal output.
    process.stdout.write(JSON.stringify(processedRecords, null, 2) + '\n');
  } catch (error) {
    // If an error occurs (e.g., invalid JSON, permission issues), 
    // we do not output anything else to stdout as per the instructions.
  }
}

// Execute the process.
processData();