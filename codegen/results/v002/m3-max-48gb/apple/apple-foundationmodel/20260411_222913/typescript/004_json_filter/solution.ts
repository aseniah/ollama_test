import fs from 'fs';
import path from 'path';

// Function to read and process the JSON data
async function processData(filePath: string): Promise<any> {
    try {
        // Read the file content
        const data = await fs.promises.readFile(filePath, 'utf8');

        // Parse the JSON data
        const jsonArray = JSON.parse(data);

        // Filter records where 'active' is true and 'age' is 30 or older
        const filteredRecords = jsonArray
            .filter(record => record.active && record.age >= 30)
            .sort((a, b) => a.name.localeCompare(b.name));

        // Return the filtered and sorted records
        return filteredRecords;
    } catch (error) {
        console.error('Error processing data:', error);
        return null;
    }
}

// Main function to execute the process
async function main(): Promise<void> {
    const filePath = path.join(__dirname, 'input', 'data.json');
    const result = await processData(filePath);

    if (result) {
        // Convert the result back to JSON string and print it
        console.log(JSON.stringify(result, null, 2));
    }
}

// Execute the main function
main().catch(console.error);