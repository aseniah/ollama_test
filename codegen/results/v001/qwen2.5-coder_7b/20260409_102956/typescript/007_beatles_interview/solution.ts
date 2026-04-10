import * as fs from 'fs';
import * as path from 'path';

const inputCsvPath = path.resolve('input', 'input.csv');
const expectedJsonPath = path.resolve('input', 'expected_format.json');

// Read the expected format to understand the transformation rules and output structure
const expectedFormat = JSON.parse(fs.readFileSync(expectedJsonPath, 'utf-8'));

// Function to calculate age as of July 1, 2025
function calculateAge(birthdate: string): number {
    const birthDateObj = new Date(birthdate);
    const targetDate = new Date(2025, 6, 1); // 0 for January
    let age = targetDate.getFullYear() - birthDateObj.getFullYear();
    if (targetDate.getMonth() < birthDateObj.getMonth() || 
        (targetDate.getMonth() === birthDateObj.getMonth() && targetDate.getDate() < birthDateObj.getDate())) {
        age--;
    }
    return age;
}

// Read and process the CSV file
const inputCsvContent = fs.readFileSync(inputCsvPath, 'utf-8');
const rows = inputCsvContent.split('\n').filter(row => row.trim() !== '');

// Transform the CSV data to match the expected JSON format
const transformedData: any[] = [];
for (let i = 1; i < rows.length; i++) {
    const columns = rows[i].split(',');
    const age = calculateAge(columns[2]); // Assuming birthdate is in the third column

    const person = {
        name: columns[0],
        email: columns[1],
        age: age
    };

    transformedData.push(person);
}

// Output the transformed data as a JSON array
console.log(JSON.stringify(transformedData, null, 2));