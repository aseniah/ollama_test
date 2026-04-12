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
 * Main function to execute the filtering and sorting logic.
 */
function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Read the file content from the specified path
    if (!fs.existsSync(filePath)) {
      // If file doesn't exist, we exit silently or handle as error. 
      // Per instructions, we just need to output the JSON array.
      return;
    }

    const fileContent = fs.readFileSync(filePath, 'utf8');
    const data: Record[] = JSON.parse(fileContent);

    // 1. Filter records where active is true AND age is 30 or older.
    // 2. Sort the resulting array by name in ascending order.
    const processedData = data
      .filter((item) => item.active === true && item.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the resulting JSON array to stdout.
    process.stdout.write(JSON.stringify(processedData, null, 2) + '\n');
  } catch (error) {
    // In a real-world production environment, we would log error details.
    // For this task, we ensure no extra output is printed to stdout.
    process.exit(1);
  }
}

main();