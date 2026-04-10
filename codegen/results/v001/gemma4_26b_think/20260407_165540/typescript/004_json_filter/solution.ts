import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input records.
 */
interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main function to process the JSON data.
 */
function processData(): void {
  // Define the path to the input file as specified.
  const inputPath = path.resolve(process.cwd(), 'input', 'data.json');

  try {
    // Check if the file exists before attempting to read.
    if (!fs.existsSync(inputPath)) {
      return;
    }

    // Read the file content synchronously.
    const fileContent = fs.readFileSync(inputPath, 'utf8');
    
    // Parse the JSON array.
    const data: Record[] = JSON.parse(fileContent);

    // Apply filters: 
    // 1. active must be true.
    // 2. age must be 30 or older.
    const filteredData = data.filter((item) => item.active === true && item.age >= 30);

    // Sort the resulting array by name in ascending order.
    filteredData.sort((a, b) => a.name.localeCompare(b.name));

    // Output the final JSON array to stdout.
    // Using JSON.stringify without indentation to ensure no extra whitespace or newlines are added.
    process.stdout.write(JSON.stringify(filteredData) + '\n');
  } catch (error) {
    // In case of a parsing error or file access error, we exit without printing to stdout
    // to comply with the "Do not output anything else" requirement.
    process.exit(1);
  }
}

// Execute the processing function.
processData();