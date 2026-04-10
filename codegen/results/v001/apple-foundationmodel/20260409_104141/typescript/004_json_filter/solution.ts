import fs from 'fs';
import path from 'path';
import { readFileSync } from 'fs';

async function filterData(inputFile: string, outputFile: string): Promise<void> {
    try {
        // Read the data from the input file
        const data = await readFileSync(inputFile, 'utf8');

        // Parse the JSON data
        const jsonArray = JSON.parse(data);

        // Filter records where active is true and age is 30 or older
        const filteredData = jsonArray.filter(record => record.active && record.age >= 30);

        // Sort the filtered records by name in ascending order
        const sortedData = filteredData.sort((a, b) => a.name.localeCompare(b.name));

        // Convert the sorted data back to JSON string
        const jsonOutput = JSON.stringify(sortedData, null, 2);

        // Write the output to the specified output file
        fs.writeFileSync(outputFile, jsonOutput);

        console.log('Filtered and sorted data written to output.json');

    } catch (error) {
        console.error('Error processing data:', error);
    }
}

if (process.argv.length < 3) {
    console.error('Usage: tsx input/data.json output/output.json');
    process.exit(1);
}

const inputFile = process.argv[2];
const outputFile = path.join(process.cwd(), 'output', 'output.json');

filterData(inputFile, outputFile);