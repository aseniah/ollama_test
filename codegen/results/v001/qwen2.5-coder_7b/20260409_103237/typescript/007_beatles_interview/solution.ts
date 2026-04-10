import * as fs from 'fs';
import * as path from 'path';

const inputFilePath = path.join(__dirname, 'input', 'input.csv');
const expectedFormatPath = path.join(__dirname, 'input', 'expected_format.json');

// Read the expected format to understand the output structure
const expectedFormat = JSON.parse(fs.readFileSync(expectedFormatPath, 'utf-8'));

// Function to calculate age as of July 1, 2025
function calculateAge(birthDate: string): number {
    const birthYear = parseInt(birthDate.split('-')[0], 10);
    return 2025 - birthYear;
}

// Read the input CSV file
const inputContent = fs.readFileSync(inputFilePath, 'utf-8');
const lines = inputContent.trim().split('\n');

// Parse the headers and map them to expected keys
const headers = lines[0].split(',');

// Function to transform each row of data
function transformRow(row: string[]): any {
    const values = row.map((value, index) => {
        if (headers[index] === 'dob') {
            return calculateAge(value);
        } else if (expectedFormat[headers[index]]) {
            return expectedFormat[headers[index]](value);
        }
        return value;
    });
    return Object.fromEntries(headers.map((header, index) => [expectedFormat[header] || header, values[index]]));
}

// Transform the data
const transformedData = lines.slice(1).map(transformRow);

// Output the JSON array
console.log(JSON.stringify(transformedData, null, 2));