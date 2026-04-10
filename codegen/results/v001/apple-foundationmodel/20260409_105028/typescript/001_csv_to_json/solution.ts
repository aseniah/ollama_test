import fs from 'fs';
import path from 'path';

async function processCsvToJson() {
  const inputPath = path.resolve(__dirname, 'input', 'data.csv');
  const outputJsonPath = path.resolve(__dirname, 'output.json');

  try {
    const csvContent = fs.readFileSync(inputPath, 'utf8');
    const rows = csvContent.split('\n');

    // Skip header
    const rowsWithoutHeader = rows.slice(1);

    const jsonArray: string[] = rowsWithoutHeader.map(row => {
      const fields = row.split(',');
      const { name, age, email, score } = fields;
      return {
        Name: name,
        Age: parseInt(age, 10),
        Email: email,
        Score: parseFloat(score)
      };
    });

    // Output JSON array
    console.log(JSON.stringify(jsonArray, null, 2));

  } catch (error) {
    console.error('Error processing CSV file:', error);
  }
}

processCsvToJson();