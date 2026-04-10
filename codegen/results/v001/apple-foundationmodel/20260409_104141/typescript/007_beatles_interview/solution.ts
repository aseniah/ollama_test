import fs from 'fs';
import path from 'path';

const inputFilePath = path.join(__dirname, 'input', 'input.csv');
const outputFilePath = path.join(__dirname, 'output', 'output.json');

// Function to calculate age as of July 1, 2025
function calculateAge(dateStr: string): number {
    const date = new Date(dateStr);
    const age = new Date('2025-07-01').getFullYear() - date.getFullYear();
    return age;
}

// Function to process each row in the CSV file
function processRow(row: string): any {
    const columns = row.split(',');
    const firstName = columns[0];
    const lastName = columns[1];
    const birthDate = columns[2];
    const age = calculateAge(birthDate);

    return {
        firstName,
        lastName,
        age
    };
}

// Read CSV file and process each row
try {
    const data = fs.readFileSync(inputFilePath, 'utf8');
    const lines = data.trim().split('\n');

    if (lines.length < 2) {
        throw new Error('CSV file is empty or improperly formatted.');
    }

    const jsonArray = lines.slice(1).map(processRow);

    // Write the JSON array to the output file
    fs.writeFileSync(outputFilePath, JSON.stringify(jsonArray, null, 2));

    console.log(JSON.stringify(jsonArray, null, 2));
} catch (error) {
    console.error('Error processing file:', error);
}