import * as fs from 'fs';

function main() {
    const csvData = fs.readFileSync('input/input.csv', 'utf-8');
    const lines = csvData.trim().split(/\r?\n/);
    const rows = lines.slice(1);

    const result: any[] = [];

    for (const line of rows) {
        const parts = line.split(',');
        if (parts.length < 7) continue;

        const nameStr = parts[0].trim();
        const birthdayStr = parts[1].trim();
        const diedStr = parts[2].trim();
        const fatherStr = parts[3].trim();
        const motherStr = parts[4].trim();
        const brotherStr = parts[5].trim();
        const sisterStr = parts[6].trim();

        const nameParts = nameStr.split(/\s+/);
        const firstName = nameParts[0];
        const lastName = nameParts[nameParts.length - 1];

        const [bMonth, bDay, bYear] = birthdayStr.split('/').map(Number);
        const birthdayFormatted = `${bYear}-${String(bMonth).padStart(2, '0')}-${String(bDay).padStart(2, '0')}`;

        // Determine reference date for age calculation
        let refYear = 2025, refMonth = 7, refDay = 1;
        if (diedStr !== 'null') {
            const [dMonth, dDay, dYear] = diedStr.split('/').map(Number);
            refYear = dYear;
            refMonth = dMonth;
            refDay = dDay;
        }

        // Calculate age
        let age = refYear - bYear;
        if (refMonth < bMonth || (refMonth === bMonth && refDay < bDay)) {
            age -= 1;
        }

        // Build relatives array
        const relatives: any[] = [];
        const addRelative = (name: string, rel: string) => {
            if (name !== 'null' && name.trim() !== '') {
                const nParts = name.split(/\s+/);
                relatives.push({
                    FirstName: nParts[0],
                    LastName: nParts[nParts.length - 1],
                    Relationship: rel
                });
            }
        };

        addRelative(fatherStr, 'Father');
        addRelative(motherStr, 'Mother');
        addRelative(brotherStr, 'Brother');
        addRelative(sisterStr, 'Sister');

        result.push({
            FirstName: firstName,
            LastName: lastName,
            Birthday: birthdayFormatted,
            Age: age,
            Relatives: relatives
        });
    }

    console.log(JSON.stringify(result, null, 2));
}

main();