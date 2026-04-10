import * as fs from 'fs';
import * as path from 'path';

// Read the files
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const expectedContent = fs.readFileSync(path.join('input', 'expected_format.json'), 'utf-8');

// Parse the expected format to understand structure
const expectedFormat = JSON.parse(expectedContent);

// Parse CSV content
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');

// Calculate age as of July 1, 2025
function calculateAge(birthDate: string): number {
    const targetDate = new Date('2025-07-01');
    const birth = new Date(birthDate);
    let age = targetDate.getFullYear() - birth.getFullYear();
    
    // Adjust if birthday hasn't occurred yet in target year
    const monthDay = targetDate.getMonth() * 100 + targetDate.getDate();
    const birthMonthDay = birth.getMonth() * 100 + birth.getDate();
    
    if (birthMonthDay > monthDay) {
        age--;
    }
    
    return age;
}

// Transform the data based on expected format
const result: any[] = [];

for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    const row: Record<string, string> = {};
    
    for (let j = 0; j < headers.length; j++) {
        row[headers[j]] = values[j] || '';
    }
    
    // Create output object based on expected format structure
    const output: any = {
        id: parseInt(row['id']) || row['id'],
        name: row['name'] || '',
        age: calculateAge(row['birth_date'] || row['dob'] || row['birthDate'] || ''),
        email: row['email'] || '',
        department: row['department'] || ''
    };
    
    result.push(output);
}

// Output JSON array to stdout
console.log(JSON.stringify(result, null, 2));