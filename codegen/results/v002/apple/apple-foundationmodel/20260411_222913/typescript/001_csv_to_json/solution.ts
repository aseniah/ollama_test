import fs from 'fs';
import path from 'path';

// Function to read and parse CSV into JSON array
async function parseCSVToJSON() {
  const csvFilePath = path.join(__dirname, 'input', 'data.csv');

  try {
    // Read the CSV file content
    const csvContent = fs.readFileSync(csvFilePath, 'utf8');

    // Split the content into lines
    const lines = csvContent.trim().split('\n');

    // Skip the header line
    const dataArray: any[] = lines.slice(1).map((line) => {
      const fields = line.split(',');
      if (fields.length === 4) {
        const name = fields[0];
        const age = parseInt(fields[1], 10);
        const email = fields[2];
        const score = parseFloat(fields[3]);
        return {
          Name: name,
          Age: age,
          Email: email,
          Score: score,
        };
      }
      return null; // Return null if parsing fails (shouldn't happen with valid input)
    });

    // Output the data as JSON
    console.log(JSON.stringify(dataArray, null, 2));

  } catch (error) {
    console.error('Error reading CSV file:', error);
  }
}

// Execute the function
parseCSVToJSON();