import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input data objects.
 */
interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main function to execute the processing logic.
 */
function main() {
  const inputFilePath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if the file exists before attempting to read
    if (!fs.existsSync(inputFilePath)) {
      // If the file doesn't exist, we exit silently as per instructions
      // regarding no extra output, though in a real scenario we might log an error.
      process.exit(0);
    }

    // Read the file content
    const fileContent = fs.readFileSync(inputFilePath, 'utf-8');
    
    // Parse the JSON content
    const data: Record[] = JSON.parse(fileContent);

    // Filter records: active must be true AND age must be 30 or older
    const filteredData = data.filter((item) => item.active === true && item.age >= 30);

    // Sort the filtered records by name in ascending order
    const sortedData = filteredData.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });

    // Output the result as a JSON array string to stdout
    process.stdout.write(JSON.stringify(sortedData, null, 2) + '\n');

  } catch (error) {
    // If an error occurs (like invalid JSON), we exit.
    // The prompt specifies "Do not output anything else", so we avoid error logging.
    process.exit(1);
  }
}

// Execute the program
main();