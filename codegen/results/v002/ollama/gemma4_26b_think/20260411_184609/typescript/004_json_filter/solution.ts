import * as fs from 'fs';

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
 * Main function to solve the task.
 * 1. Reads input from 'input/data.json'.
 * 2. Filters records where 'active' is true and 'age' is 30 or older.
 * 3. Sorts the resulting records by 'name' in ascending order.
 * 4. Outputs the resulting JSON array to stdout.
 */
function main() {
  const inputPath = 'input/data.json';

  try {
    // Verify the existence of the file to avoid throwing unhandled exceptions
    if (fs.existsSync(inputPath)) {
      // Read the file content as a UTF-8 string
      const fileContent = fs.readFileSync(inputPath, 'utf8');
      
      // Parse the JSON string into an array of DataRecord objects
      const records: DataRecord[] = JSON.parse(fileContent);

      // Filter the records based on the criteria:
      // - active must be true
      // - age must be 30 or greater
      const filteredRecords = records.filter((record) => {
        return record.active === true && record.age >= 30;
      });

      // Sort the filtered records by name in ascending order using localeCompare
      // localeCompare is used for robust string comparison
      filteredRecords.sort((a, b) => a.name.localeCompare(b.name));

      // Convert the resulting array back to a JSON string and write to stdout.
      // We do not use indentation (pretty-printing) to ensure the output 
      // is a clean, single-line JSON array unless required.
      process.stdout.write(JSON.stringify(filteredRecords));
    }
  } catch (error) {
    // The requirement specifies: "Do not output anything else."
    // Therefore, we catch errors and prevent error messages from being sent to stdout.
  }
}

// Execute the program
main();