import fs from 'fs';
import path from 'path';

// Function to calculate age in years
function calculateAge(birthday: string): number {
    const date = new Date(birthday);
    const currentDate = new Date(new Date().getFullYear(), 6, 1); // July 1, 2025
    return currentDate.getFullYear() - date.getFullYear() - ((currentDate.getMonth() + 1) - date.getMonth() - 1);
}

async function main() {
    const inputFilePath = path.join(__dirname, 'input', 'input.csv');
    const outputFilePath = path.join(__dirname, 'output', 'output.json');

    // Read the CSV file
    const csvData = fs.readFileSync(inputFilePath, 'utf8').trim();
    const rows = csvData.split('\n').filter(row => !row.trim().startsWith('#'));

    // Parse each row and calculate age
    const data: any[] = [];
    for (const row of rows) {
        const fields = row.split(',').map(field => field.trim());
        if (fields.length !== 7) continue; // Ensure all fields are present

        const { Name, Birthday, Died, Father, Mother, Brother, Sister } = fields;
        
        // Calculate age
        const age = calculateAge(Birthday);

        // Prepare JSON object
        const person: any = {
            Name,
            Birthday,
            Age: age,
            Relatives: [
                { FirstName: Father, LastName: Father, Relationship: 'Father' },
                { FirstName: Mother, LastName: Mother, Relationship: 'Mother' },
                { FirstName: Brother, LastName: Brother, Relationship: 'Brother' },
                { FirstName: Sister, LastName: Sister, Relationship: 'Sister' }
            ]
        };

        data.push(person);
    }

    // Write JSON to output file
    fs.writeFileSync(outputFilePath, JSON.stringify(data, null, 2));
}

main();