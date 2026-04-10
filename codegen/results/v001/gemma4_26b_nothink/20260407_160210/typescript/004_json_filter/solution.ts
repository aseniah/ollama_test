import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of the input objects
 */
interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

/**
 * Main execution function
 */
function main() {
  const filePath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if file exists before attempting to read
    if (!fs.existsSync(filePath)) {
      process.exit(0);
    }

    // Read file content
    const rawData = fs.readFileSync(filePath, 'utf8');
    
    // Parse JSON content
    const data: Record[] = JSON.parse(rawData);

    // Filter: active is true AND age is 30 or older
    // Sort: by name ascending
    const filteredData = data
      .filter((item) => item.active === true && item.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the result to stdout as a JSON string
    process.stdout.write(JSON.stringify(filteredData, null, 2) + '\n');
  } catch (error) {
    // In a production environment, we might log to stderr.
    // Per requirements, we output nothing else to stdout.
    process.exit(1);
  }
}

main();