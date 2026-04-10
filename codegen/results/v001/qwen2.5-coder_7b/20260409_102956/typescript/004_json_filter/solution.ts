import * as fs from 'fs';
import * as path from 'path';

// Read command-line arguments
const [,, inputPath] = process.argv;

// Ensure the file exists and is a valid JSON array of objects with specified fields
if (!inputPath) {
  console.error('Please provide the path to the input JSON file.');
  process.exit(1);
}

try {
  // Read the file content
  const fileContent = fs.readFileSync(inputPath, 'utf8');
  
  // Parse the JSON content
  const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(fileContent);

  // Filter and sort the data
  const filteredData = data
    .filter(record => record.active && record.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  // Output the filtered and sorted data as a JSON array to stdout
  console.log(JSON.stringify(filteredData, null, 2));
} catch (error) {
  if (error instanceof SyntaxError) {
    console.error('Invalid JSON format in the input file.');
  } else {
    console.error('An error occurred while reading or processing the file:', error.message);
  }
  process.exit(1);
}