import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input JSON objects.
 */
interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main function to process the data.
 */
function main() {
  // Define the path to the input file
  const inputPath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if the file exists
    if (!fs.existsSync(inputPath)) {
      // If file doesn't exist, we exit silently as per instructions
      // (Do not output anything else).
      process.exit(0);
    }

    // Read the file content
    const fileContent = fs.readFileSync(inputPath, 'utf8');

    // Parse the JSON content
    const data: Record[] = JSON.parse(fileContent);

    // Filter records: active must be true AND age must be 30 or older
    const filteredData = data.filter((item) => item.active === true && item.age >= 30);

    // Sort the filtered records by name ascending
    filteredData.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });

    // Output the resulting JSON array to stdout
    // Using JSON.stringify with 2 spaces for clean formatting, 
    // but standard JSON stringify is fine.
    process.stdout.write(JSON.stringify(filteredData, null, 2) + '\n');

  } catch (error) {
    // In a production environment, you might log to stderr.
    // Given the strict constraint "Do not output anything else", 
    // we exit quietly on error to avoid corrupting the expected output format.
    process.exit(1);
  }
}

// Execute the script
main();