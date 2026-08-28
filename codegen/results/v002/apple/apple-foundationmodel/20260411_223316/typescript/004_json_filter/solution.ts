import fs from 'fs';
import path from 'path';

// Function to filter and sort the records
const processData = async (inputFilePath: string, outputFilePath: string): Promise<void> => {
  try {
    // Read the JSON file
    const content = await fs.readFileSync(inputFilePath, 'utf8');
    const data = JSON.parse(content);

    // Filter records where 'active' is true and 'age' is 30 or older
    const filteredData = data.filter((record) => {
      return record.active && record.age >= 30;
    });

    // Sort the filtered records by 'name' in ascending order
    const sortedData = filteredData.sort((a, b) => {
      return a.name.localeCompare(b.name);
    });

    // Write the sorted data to the output file
    await fs.writeFileSync(outputFilePath, JSON.stringify(sortedData, null, 2));
  } catch (error) {
    console.error('An error occurred:', error);
  }
};

// Main function to handle command-line arguments
async function main() {
  if (process.argv.length < 3) {
    console.error('Usage: tsx input/data.json output/filtered_data.json');
    return;
  }

  const inputFilePath = path.resolve(process.argv[2]);
  const outputFilePath = path.resolve(process.argv[3]);

  // Process the data
  await processData(inputFilePath, outputFilePath);

  console.log('Data processing completed.');
}

// Run the main function
main();