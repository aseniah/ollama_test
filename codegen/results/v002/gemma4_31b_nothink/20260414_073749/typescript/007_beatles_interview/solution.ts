import * as fs from 'fs';
import * as path from 'path';

/**
 * Calculates age as of July 1, 2025.
 * If the person died, the age is calculated as of the date of death.
 */
function calculateAge(birthdayStr: string, diedStr: string | null): number {
    const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)
    
    // Parse MM/DD/YYYY
    const parseDate = (s: string) => {
        const [m, d, y] = s.split('/').map(Number);
        return new Date(y, m - 1, d);
    };

    const birthDate = parseDate(birthdayStr);
    const endDate = diedStr ? parseDate(diedStr) : referenceDate;

    let age = endDate.getFullYear() - birthDate.getFullYear();
    const m = endDate.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && endDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

function solve() {
    const csvPath = path.join('input', 'input.csv');
    const content = fs.readFileSync(csvPath, 'utf8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    if (lines.length === 0) return;

    const headers = lines[0].split(',');
    const dataLines = lines.slice(1);

    const result = dataLines.map(line => {
        const values = line.split(',');
        const row: Record<string, string> = {};
        headers.forEach((h, i) => {
            row[h] = values[i];
        });

        const fullNames = row['Name'].split(' ');
        const firstName = fullNames[0];
        const lastName = fullNames[fullNames.length - 1];
        
        const birthday = row['Birthday'];
        const died = row['Died'] === 'null' ? null : row['Died'];
        
        const relatives = [];
        const relativeTypes = ['Father', 'Mother', 'Brother', 'Sister'];
        
        relativeTypes.forEach(type => {
            const val = row[type];
            if (val && val !== 'null') {
                const parts = val.split(' ');
                relatives.push({
                    FirstName: parts[0],
                    LastName: parts[parts.length - 1],
                    Relationship: type
                });
            }
        });

        return {
            FirstName: firstName,
            LastName: lastName,
            Birthday: `${birthday.split('/')[2]}-${birthday.split('/')[0].padStart(2, '0')}-${birthday.split('/')[1].padStart(2, '0')}`,
            Age: calculateAge(birthday, died),
            Relatives: relatives
        };
    });

    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

solve();