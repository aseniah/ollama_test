import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface defining the structure of the input records
 */
interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function solve() {
  try {
    // Define the path to the input file
    const filePath = path.join(process.cwd(), 'input', 'data.json');

    // Read the file content
    const rawData = fs.readFileSync(filePath, 'utf8');

    // Parse the JSON data
    const data: Record[] = JSON.parse(rawData);

    // Filter: records where active is true and age is 30 or older
    const filteredData = data.filter((item) => item.active === true && item.age >= 30);

    // Sort: by name ascending
    const sortedData = filteredData.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });

    // Output the resulting JSON array to stdout
    process.stdout.write(JSON.stringify(sortedData, null, 2) + '\n');
  } catch (error) {
    // In a production environment, we'd handle specific errors (e.g., FileNotFound),
    // but based on the requirements, we just need the logic for the provided input.
    process.exit(1);
  }
}

// Execute the solver
solve();