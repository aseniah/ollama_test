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
  const inputPath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if file exists before attempting to read
    if (!fs.existsSync(inputPath)) {
      process.exit(0);
    }

    // Read file content
    const rawData = fs.readFileSync(inputPath, 'utf8');
    
    // Parse JSON array
    const records: Record[] = JSON.parse(rawData);

    // Filter: active is true AND age is 30 or older
    // Sort: name ascending
    const result = records
      .filter((record) => record.active === true && record.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    // Output the resulting JSON array to stdout
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
  } catch (error) {
    // In case of error (e.g. malformed JSON), exit without extra output 
    // to adhere to the constraint: "Do not output anything else."
    process.exit(1);
  }
}

// Execute the processing logic
processData();