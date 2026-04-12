import fs from 'fs';
import path from 'path';

// Function to process the JSON data
async function processData(filePath: string): Promise<any> {
  const filePathWithoutExtension = path.basename(filePath).split('.')[0];
  const jsonData = await fs.promises.readFile(filePath, 'utf8');
  const data = JSON.parse(jsonData);

  // Filter records where active is true and age is 30 or older
  const filteredData = data.filter(record => {
    return record.active && record.age >= 30;
  });

  // Sort the filtered data by name in ascending order
  const sortedData = filteredData.sort((a, b) => {
    return a.name.localeCompare(b.name);
  });

  return sortedData;
}

// Main function to handle command-line arguments
async function main(): Promise<void> {
  if (process.argv.length !== 2) {
    console.error('Usage: tsx input/data.json');
    return;
  }

  const filePath = process.argv[2];
  try {
    const result = await processData(filePath);
    console.log(JSON.stringify(result, null, 2));
  } catch (error) {
    console.error('Error processing data:', error);
  }
}

// Execute the main function
main();