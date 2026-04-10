import fs from 'fs';
import path from 'path';

function calculateAges(birthDate: string): number {
    const dateComponents = birthDate.split('-').map(Number);
    const year = dateComponents[2];
    const age = 2025 - year;
    return age;
}

async function processCSVFile(filePath: string): Promise<string> {
    const fileContents = await fs.promises.readFile(filePath, 'utf8');
    const lines = fileContents.trim().split('\n');
    
    const jsonArray: any[] = [];

    lines.forEach(line => {
        const fields = line.split(',');
        if (fields.length < 2) {
            console.warn(`Skipping malformed line: ${line}`);
            return;
        }
        const firstName = fields[0].trim();
        const lastName = fields[1].trim();
        const birthDate = fields[2].trim();

        // Calculate age
        const age = calculateAges(birthDate);
        
        // Create object with calculated age
        const person = {
            firstName,
            lastName,
            age
        };

        jsonArray.push(person);
    });

    return JSON.stringify(jsonArray, null, 2);
}

async function main() {
    const inputFilePath = path.join(__dirname, 'input', 'input.csv');
    const outputFilePath = path.join(__dirname, 'output', 'output.json');

    try {
        const output = await processCSVFile(inputFilePath);
        console.log(output);
    } catch (error) {
        console.error('Error processing file:', error);
    }
}

main().catch(console.error);