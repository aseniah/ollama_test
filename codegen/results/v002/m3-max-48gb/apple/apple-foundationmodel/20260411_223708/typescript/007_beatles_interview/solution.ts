const fs = require('fs');
const path = require('path');

function calculateAge(birthday: string, currentDate: Date) {
    const dateFormatter = new Date(birthday);
    const ageComponents = currentDate.getFullYear() - dateFormatter.getFullYear();
    const monthDiff = currentDate.getMonth() - dateFormatter.getMonth();
    const dayDiff = currentDate.getDate() - dateFormatter.getDate();
    
    if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
        ageComponents--;
    }
    
    return ageComponents;
}

async function main() {
    const inputFilePath = path.join(__dirname, 'input/input.csv');
    const expectedFormatFilePath = path.join(__dirname, 'input/expected_format.json');

    // Read the CSV file and parse it into an array of objects
    const csvData = fs.readFileSync(inputFilePath, 'utf8').toString().trim().split('\n');
    const csvRows = csvData.map(row => row.split(',').map(cell => cell.trim()));
    
    // Ensure the CSV file has the expected columns
    const expectedColumns = ['Name', 'Birthday', 'Died', 'Father', 'Mother', 'Brother', 'Sister'];
    if (!expectedColumns.every(column => csvRows[0].includes(column))) {
        throw new Error('CSV file does not have the expected columns.');
    }

    const people: any[] = [];

    csvRows.forEach(row => {
        const person: any = {
            FirstName: row[0],
            LastName: row[1],
            Birthday: row[2],
            Age: calculateAge(row[2], new Date()),
            Relatives: []
        };

        if (row.length > 2) {
            const father = row[3] === 'null' ? null : row[3];
            const mother = row[4] === 'null' ? null : row[4];
            const brother = row[5] === 'null' ? null : row[5];
            const sister = row[6] === 'null' ? null : row[6];

            person.Relatives.push({
                FirstName: brother || '',
                LastName: brother || '',
                Relationship: brother || '',
                ...(mother && { Relationship: 'Mother' }),
                ...(father && { Relationship: 'Father' }),
                ...(sister && { Relationship: 'Sister' })
            });
        }

        people.push(person);
    });

    // Convert people into a JSON array
    const jsonArray = JSON.stringify(people, null, 2);
    console.log(jsonArray);
}

main();