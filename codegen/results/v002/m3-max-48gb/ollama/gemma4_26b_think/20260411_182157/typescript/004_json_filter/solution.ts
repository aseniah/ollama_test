import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input JSON objects.
 */
interface DataRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main function to execute the task logic.
 * 1. Reads the file input/data.json.
 * 2. Filters records where active is true and age is 30 or older.
 * 3. Sorts the remaining records by name in ascending order.
 * 4. Outputs the resulting JSON array to stdout.
 */
function main() {
  // The prompt specifies the file is located at input/data.json.
  // We resolve this path relative to the current working directory.
  const filePath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if the file exists before attempting to read it.
    if (!fs.existsSync(filePath)) {
      return;
    }

    // Read file content synchronously.
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    
    // Parse the JSON content.
    const data: DataRecord[] = JSON.parse(fileContent);

    // Perform filtering: active must be true and age must be >= 30.
    // Perform sorting: sort by 'name' property in ascending order.
    const processedData = data
      .filter((record) => record.active === true && record.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Convert the result back to a JSON string and print to stdout.
    // We use process.stdout.write to ensure no extra characters are added 
    // beyond the JSON string and a single newline.
    process.stdout.write(JSON.stringify(processedData) + '\n');
  } catch (error) {
    // If any error occurs (e.g., invalid JSON), we exit silently 
    // to comply with the instruction "Do not output anything else".
  }
}

// Execute the program.
main();