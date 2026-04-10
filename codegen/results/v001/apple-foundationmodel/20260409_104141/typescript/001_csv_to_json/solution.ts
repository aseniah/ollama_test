import fs from 'fs';
import path from 'path';

async function processCSVFile(filePath: string) {
  try {
    // Read the file content
    const fileContent = fs.readFileSync(filePath, 'utf8');

    // Split the content by lines
    const lines = fileContent.split('\n');

    // The header line is the first line
    const headerLine = lines[0];

    // Process each subsequent line (excluding the header)
    const jsonArray: any[] = lines.slice(1).map((line) => {
      const fields = line.split(',');
      if (fields.length < 4) {
        throw new Error('Invalid CSV format');
      }
      // Ensure fields are strings and numbers as expected
      const { name, age, email, score } = {
        name: fields[0].trim(),
        age: Number(fields[1].trim()),
        email: fields[2].trim(),
        score: Number(fields[3].trim()),
      };
      return {
        Name: name,
        Age: age,
        Email: email,
        Score: score,
      };
    });

    // Output the JSON array
    console.log(JSON.stringify(jsonArray, null, 2));
  } catch (error) {
    console.error('Error processing CSV file:', error);
  }
}

// Example usage
const inputFilePath = process.argv[2];
processCSVFile(inputFilePath);