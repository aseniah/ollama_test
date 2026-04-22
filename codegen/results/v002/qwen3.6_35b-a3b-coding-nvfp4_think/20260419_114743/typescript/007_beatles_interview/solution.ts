import * as fs from 'fs';

const inputPath = 'input/input.csv';
const csvContent = fs.readFileSync(inputPath, 'utf-8').trim();
const lines = csvContent.split('\n').slice(1);

const result: any[] = [];

for (const line of lines) {
    const cols = line.split(',').map(c => c.trim());
    const nameParts = cols[0].split(/\s+/);
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];
    
    const birthdayStr = cols[1];
    const [bMonthStr, bDayStr, bYearStr] = birthdayStr.split('/');
    const bMonth = parseInt(bMonthStr, 10);
    const bDay = parseInt(bDayStr, 10);
    const bYear = parseInt(bYearStr, 10);
    
    const birthdayFormatted = `${bYear}-${String(bMonth).padStart(2, '0')}-${String(bDay).padStart(2, '0')}`;
    
    const refYear = 2025;
    const refMonth = 7;
    const refDay = 1;
    
    let age = refYear - bYear;
    if (refMonth < bMonth || (refMonth === bMonth && refDay < bDay)) {
        age -= 1;
    }
    
    const relatives: any[] = [];
    const relationships = ['Father', 'Mother', 'Brother', 'Sister'];
    
    for (let i = 0; i < relationships.length; i++) {
        const relativeName = cols[3 + i];
        if (relativeName !== 'null' && relativeName !== '') {
            const relParts = relativeName.split(/\s+/);
            relatives.push({
                FirstName: relParts[0],
                LastName: relParts[relParts.length - 1],
                Relationship: relationships[i]
            });
        }
    }
    
    result.push({
        FirstName: firstName,
        LastName: lastName,
        Birthday: birthdayFormatted,
        Age: age,
        Relatives: relatives
    });
}

console.log(JSON.stringify(result, null, 2));