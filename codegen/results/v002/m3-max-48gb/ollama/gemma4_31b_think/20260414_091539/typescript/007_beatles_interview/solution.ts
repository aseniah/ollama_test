import * as fs from 'fs';
import * as path from 'path';

/**
 * Helper to convert date from M/D/YYYY to YYYY-MM-DD and a Date object.
 */
function parseDate(dateStr: string): { iso: string, date: Date } {
    const parts = dateStr.split('/');
    const month = parts[0].padStart(2, '0');
    const day = parts[1].padStart(2, '0');
    const year = parts[2];
    const iso = `${year}-${month}-${day}`;
    return { iso, date: new Date(`${iso}`) };
}

/**
 * Calculate age based on a birth date and a reference date.
 */
function calculateAge(birthdayStr: string, referenceDateStr: string | null): number {
    const birth = parseDate(birthdayStr);
    const ref = referenceDateStr ? parseDate(referenceDateStr) : { 
        iso: '2025-07-01', 
        date: new Date('2025-07-01') 
    };

    const bDate = birth.date;
    const rDate = ref.date;

    let age = rDate.getFullYear() - bDate.getFullYear();
    const mDiff = rDate.getMonth() - bDate.getMonth();
    const dDiff = rDate.getDate() - bDate.getDate();

    if (mDiff < 0 || (mDiff === 0 && dDiff < 0)) {
        age--;
    }

    return age;
}

/**
 * Splits a name into FirstName and LastName.
 * Rule inferred from expected_format: First word is FirstName, last word is LastName.
 */
function splitName(fullName: string): { firstName: string, lastName: string } {
    const parts = fullName.trim().split(/\s+/);
    return {
        firstName: parts[0],
        lastName: parts[parts.length - 1]
    };
}

function main() {
    const inputPath = path.join('input', 'input.csv');
    if (!fs.existsSync(inputPath)) {
        process.exit(1);
    }

    const content = fs.readFileSync(inputPath, 'utf8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    if (lines.length === 0) {
        console.log('[]');
        return;
    }

    const header = lines[0].split(',');
    const dataLines = lines.slice(1);

    const result = dataLines.map(line => {
        const values = line.split(',');
        const row: any = {};
        header.forEach((key, index) => {
            row[key] = values[index];
        });

        const nameInfo = splitName(row['Name']);
        const birthdayIso = parseDate(row['Birthday']).iso;
        const age = calculateAge(row['Birthday'], row['Died'] === 'null' ? null : row['Died']);

        const relatives = [];
        const relativeKeys = ['Father', 'Mother', 'Brother', 'Sister'];
        
        relativeKeys.forEach(key => {
            const relName = row[key];
            if (relName && relName !== 'null') {
                const relInfo = splitName(relName);
                relatives.push({
                    FirstName: relInfo.firstName,
                    LastName: relInfo.lastName,
                    Relationship: key
                });
            }
        });

        return {
            FirstName: nameInfo.firstName,
            LastName: nameInfo.lastName,
            Birthday: birthdayIso,
            Age: age,
            Relatives: relatives
        };
    });

    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();