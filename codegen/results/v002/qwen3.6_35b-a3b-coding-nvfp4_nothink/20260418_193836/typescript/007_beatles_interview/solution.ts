import * as fs from 'fs';
import * as path from 'path';

// Read the input CSV file
const csvData = fs.readFileSync(path.join(__dirname, 'input', 'input.csv'), 'utf-8');
const lines = csvData.trim().split('\n');

// Parse the header
const headers = lines[0].split(',');

// Reference date for age calculation
const referenceDate = new Date('2025-07-01');

const result: any[] = [];

for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;
    
    const values = line.split(',');
    
    // Map values to headers
    const record: { [key: string]: string } = {};
    headers.forEach((header, index) => {
        record[header] = values[index] ? values[index].trim() : '';
    });
    
    // Parse name into first and last name
    const fullName = record['Name'];
    const nameParts = fullName.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];
    
    // Parse birthday
    const birthdayStr = record['Birthday'];
    const [month, day, year] = birthdayStr.split('/').map(Number);
    const birthday = new Date(year, month - 1, day); // month is 0-indexed
    
    // Calculate age as of July 1, 2025
    let age = referenceDate.getFullYear() - birthday.getFullYear();
    const monthDiff = referenceDate.getMonth() - birthday.getMonth();
    if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthday.getDate())) {
        age--;
    }
    
    // Build relatives array
    const relatives: any[] = [];
    
    const father = record['Father'];
    if (father && father !== 'null') {
        const fatherParts = father.split(' ');
        relatives.push({
            FirstName: fatherParts[0],
            LastName: fatherParts[fatherParts.length - 1],
            Relationship: 'Father'
        });
    }
    
    const mother = record['Mother'];
    if (mother && mother !== 'null') {
        const motherParts = mother.split(' ');
        relatives.push({
            FirstName: motherParts[0],
            LastName: motherParts[motherParts.length - 1],
            Relationship: 'Mother'
        });
    }
    
    const brother = record['Brother'];
    if (brother && brother !== 'null') {
        const brotherParts = brother.split(' ');
        relatives.push({
            FirstName: brotherParts[0],
            LastName: brotherParts[brotherParts.length - 1],
            Relationship: 'Brother'
        });
    }
    
    const sister = record['Sister'];
    if (sister && sister !== 'null') {
        const sisterParts = sister.split(' ');
        relatives.push({
            FirstName: sisterParts[0],
            LastName: sisterParts[sisterParts.length - 1],
            Relationship: 'Sister'
        });
    }
    
    result.push({
        FirstName: firstName,
        LastName: lastName,
        Birthday: birthdayStr,
        Age: age,
        Relatives: relatives
    });
}

// Output the result as JSON
console.log(JSON.stringify(result, null, 2));