import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface representing the structure of each record in the input JSON.
 */
interface UserRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  try {
    // Define the path to the input file
    const filePath = path.join('input', 'data.json');
    
    // Read the file content synchronously
    const fileContent = fs.readFileSync(filePath, 'utf8');
    
    // Parse the JSON string into an array of UserRecord objects
    const data: UserRecord[] = JSON.parse(fileContent);
    
    // Filter records: active must be true AND age must be 30 or older
    const filteredData = data.filter(record => record.active === true && record.age >= 30);
    
    // Sort records by name in ascending order
    filteredData.sort((a, b) => a.name.localeCompare(b.name));
    
    // Output the resulting array as a JSON string to stdout
    process.stdout.write(JSON.stringify(filteredData, null, 2) + '\n');
  } catch (error) {
    // In a production environment, you might handle specific errors (e.g., file not found)
    // For this task, we ensure it doesn't output anything other than the requested JSON
    process.exit(1);
  }
}

// Execute the main function
main();