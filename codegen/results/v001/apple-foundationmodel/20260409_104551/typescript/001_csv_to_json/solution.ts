import fs from 'fs';

async function processCsvToJson() {
  try {
    // Read the CSV file content
    const csvContent = await fs.promises.readFile('input/data.csv', 'utf8');

    // Split the content into rows
    const rows = csvContent.split('\n');

    // Skip the header
    const dataArray: any[] = rows.slice(1).map(row => {
      const columns = row.split(',');
      const name = columns[0].trim();
      const age = Number(columns[1].trim());
      const email = columns[2].trim();
      const score = Number(columns[3].trim());
      return { name, age, email, score };
    });

    // Output the JSON array
    console.log(JSON.stringify(dataArray, null, 2));
  } catch (error) {
    console.error('Error processing CSV file:', error);
  }
}

// Call the function to process CSV to JSON
processCsvToJson();