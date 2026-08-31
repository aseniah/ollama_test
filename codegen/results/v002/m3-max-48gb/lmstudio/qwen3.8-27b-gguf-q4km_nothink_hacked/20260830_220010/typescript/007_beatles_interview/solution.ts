import * as fs from 'fs';
import * as path from 'path';

interface Relative {
    FirstName: string;
    LastName: string;
    Relationship: string;
}

interface Person {
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Relative[];
}

function main() {
    const inputPath = path.join('input', 'input.csv');
    const content = fs.readFileSync(inputPath, 'utf-8');
    const lines = content.trim().split('\n');
    
    if (lines.length < 2) {
        console.log('[]');
        return;
    }

    const headers = lines[0].split(',');
    const people: Person[] = [];
    
    // Reference date: July 1, 2025
    const refDate = new Date('2025-07-01T00:00:00Z');
    const refYear = refDate.getUTCFullYear();
    const refMonth = refDate.getUTCMonth() + 1;
    const refDay = refDate.getUTCDate();

    for (let i = 1; i < lines.length; i++) {
        const line = lines[i];
        if (!line.trim()) continue;
        
        const fields = line.split(',').map(f => f.trim());
        
        const name = fields[0];
        const birthday = fields[1];
        const father = fields[3];
        const mother = fields[4];
        const brother = fields[5];
        const sister = fields[6];

        // Parse name - assume format "FirstName MiddleName(s) LastName"
        const nameParts = name.split(' ');
        const lastName = nameParts[nameParts.length - 1];
        const firstName = nameParts[0];

        // Parse birthday M/D/YYYY
        const bdParts = birthday.split('/');
        const bdMonth = parseInt(bdParts[0], 10);
        const bdDay = parseInt(bdParts[1], 10);
        const bdYear = parseInt(bdParts[2], 10);
        const birthdayFormatted = `${bdYear.toString().padStart(4, '0')}-${bdMonth.toString().padStart(2, '0')}-${bdDay.toString().padStart(2, '0')}`;

        // Calculate age as of July 1, 2025
        let age = refYear - bdYear;
        if (bdMonth > refMonth || (bdMonth === refMonth && bdDay > refDay)) {
            age--;
        }

        const relatives: Relative[] = [];

        function addRelative(value: string, relationship: string) {
            if (value && value.toLowerCase() !== 'null') {
                const relParts = value.split(' ');
                const relLastName = relParts[relParts.length - 1];
                const relFirstName = relParts[0];
                relatives.push({
                    FirstName: relFirstName,
                    LastName: relLastName,
                    Relationship: relationship
                });
            }
        }

        addRelative(father, 'Father');
        addRelative(mother, 'Mother');
        addRelative(brother, 'Brother');
        addRelative(sister, 'Sister');

        people.push({
            FirstName: firstName,
            LastName: lastName,
            Birthday: birthdayFormatted,
            Age: age,
            Relatives: relatives
        });
    }

    console.log(JSON.stringify(people, null, 2));
}

main();