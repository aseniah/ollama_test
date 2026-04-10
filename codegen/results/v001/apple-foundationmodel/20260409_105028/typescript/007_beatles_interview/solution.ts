import fs from 'fs';
import path from 'path';

async function main() {
    const inputFilePath = path.join(__dirname, 'input', 'input.csv');
    const expectedOutputFilePath = path.join(__dirname, 'output', 'expected_format.json');

    const inputData: string[] = await fs.readFile(inputFilePath, 'utf8').then(data => data.split('\n'));

    const csvRows = inputData.filter(row => row.trim() !== '');

    const expectedOutput: any[] = [];

    csvRows.forEach(row => {
        const columns = row.split(',');
        if (columns.length < 2) return; // Ensure there are at least two columns

        const firstName = columns[0].trim();
        const dateOfBirth = columns[1].trim();

        const age = calculateAgeAsOf(new Date(dateOfBirth), new Date(`2025-07-01`));

        expectedOutput.push({
            firstName,
            age
        });
    });

    fs.writeFileSync(expectedOutputFilePath, JSON.stringify(expectedOutput, null, 2));

    console.log(JSON.stringify(expectedOutput, null, 2));
}

async function calculateAgeAsOf(dob: Date, today: Date): Date {
    return today.getFullYear() - dob.getFullYear() - ((today.getMonth() + 1) < (dob.getMonth() + 1) ? 1 : 0);
}

main();