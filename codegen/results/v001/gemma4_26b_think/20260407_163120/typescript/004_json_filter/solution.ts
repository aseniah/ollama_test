import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input data objects.
 */
interface DataRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main function to process the JSON file.
 * Reads input/data.json, filters by criteria, sorts by name, and outputs JSON.
 */
function processData() {
  // Define the path to the required input file
  const filePath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Read the file content synchronously
    const fileContent = fs.readFileSync(filePath, 'utf8');
    
    // Parse the JSON array
    const records: DataRecord[] = JSON.parse(fileContent);

    // 1. Filter: active must be true AND age must be 30 or older
    // 2. Sort: name ascending using localeCompare for alphabetical correctness
    const result = records
      .filter((record) => record.active === true && record.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the resulting array as a JSON string to stdout
    // We use JSON.stringify without indentation to keep the output clean, 
    // though a single line or formatted JSON are both valid JSON.
    process.stdout.write(JSON.stringify(result) + '\n');
  } catch (error) {
    // The requirement states: "Do not output anything else."
    // Therefore, we catch errors (like file not found or invalid JSON) 
    // and remain silent to ensure only the valid JSON array is output.
  }
}

// Execute the logic
processData();