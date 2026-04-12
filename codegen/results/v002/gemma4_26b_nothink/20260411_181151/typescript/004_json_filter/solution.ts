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
 * Main function to process the JSON file.
 */
function main() {
  const inputPath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Read the file from the specified path
    if (!fs.existsSync(inputPath)) {
      // If file doesn't exist, we exit silently or handle error.
      // Based on requirements, we assume the file is provided.
      return;
    }

    const fileContent = fs.readFileSync(inputPath, 'utf-8');
    const data: Record[] = JSON.parse(fileContent);

    // Filter: active is true AND age is 30 or older
    // Sort: name ascending
    const result = data
      .filter((record) => record.active === true && record.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the resulting JSON array to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
  } catch (error) {
    // In a production environment, we might log to stderr.
    // For this task, we ensure no extra output is sent to stdout.
    process.exit(1);
  }
}

// Run the program
main();