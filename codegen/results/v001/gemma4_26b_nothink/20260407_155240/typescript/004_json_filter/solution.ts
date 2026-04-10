import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a JSON file from 'input/data.json',
 * filters records based on specific criteria (active is true, age >= 30),
 * and outputs the sorted result to stdout.
 */

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'data.json');

  try {
    // Check if file exists before attempting to read
    if (!fs.existsSync(inputPath)) {
      // If the file doesn't exist, we exit silently or with error 
      // as per the instruction to "not output anything else" 
      // (though in a real CLI we might want an error message).
      process.exit(1);
    }

    const rawData = fs.readFileSync(inputPath, 'utf8');
    const data: Record[] = JSON.parse(rawData);

    // Filter: active must be true AND age must be 30 or older
    const filteredData = data.filter(item => item.active === true && item.age >= 30);

    // Sort: name ascending
    const sortedData = filteredData.sort((a, b) => {
      if (a.name < b.name) return -1;
      if (a.name > b.name) return 1;
      return 0;
    });

    // Output only the JSON array to stdout
    process.stdout.write(JSON.stringify(sortedData, null, 2) + '\n');

  } catch (error) {
    // If there's a parsing error or file error, exit with error code
    process.exit(1);
  }
}

main();