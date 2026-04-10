import fs from 'fs';
import path from 'path';

async function main() {
    const inputPath = path.join(__dirname, 'input', 'data.json');
    const outputPath = path.join(__dirname, 'output', 'filtered_data.json');

    // Read the input file
    const inputData = await fs.promises.readFile(inputPath, 'utf8');

    // Parse the input data
    const jsonData = JSON.parse(inputData);

    // Filter records where 'active' is true and 'age' is 30 or older
    const filteredData = jsonData.filter(record => record.active && record.age >= 30);

    // Sort the filtered records by 'name' in ascending order
    const sortedData = filteredData.sort((a, b) => a.name.localeCompare(b.name));

    // Write the output to a new file
    await fs.promises.writeFile(outputPath, JSON.stringify(sortedData, null, 2));

    console.log('Filtered and sorted data output written to:', outputPath);
}

main().catch((error) => {
    console.error('An error occurred:', error);
});