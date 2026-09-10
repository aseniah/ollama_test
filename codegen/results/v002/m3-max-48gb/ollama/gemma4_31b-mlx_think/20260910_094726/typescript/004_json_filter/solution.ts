import * as fs from 'fs';
import * as path from 'path';

/**
 * Define the interface for the record objects based on the requirements.
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
    const inputFilePath = path.join('input', 'data.json');

    // Read the file content synchronously
    const fileContent = fs.readFileSync(inputFilePath, 'utf8');

    // Parse the JSON content into an array of UserRecord objects
    const data: UserRecord[] = JSON.parse(fileContent);

    // Filter records: active must be true and age must be 30 or older
    const filteredData = data.filter(record => record.active === true && record.age >= 30);

    // Sort records by name in ascending order
    const sortedData = filteredData.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });

    // Output the result as a JSON array to stdout
    process.stdout.write(JSON.stringify(sortedData, null, 2) + '\n');
  } catch (error) {
    // Silently fail or handle error if necessary, though the prompt asks for only the JSON output.
    process.exit(1);
  }
}

main();