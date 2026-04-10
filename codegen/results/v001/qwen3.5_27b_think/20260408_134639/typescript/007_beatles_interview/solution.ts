import * as fs from 'fs';
import * as path from 'path';

// Read and parse CSV
function readCSV(filePath: string): string[][] {
    const content = fs.readFileSync(filePath, 'utf-8');
    const lines = content.trim().split('\n');
    const headers = lines[0].split(',').map(h => h.trim());
    const rows = [];
    
    for (let i = 1; i < lines.length; i++) {
        const line = lines[i];
        if (!line.trim()) continue;
        
        // Handle potential quoted fields
        const values: string[] = [];
        let current = '';
        let inQuotes = false;
        
        for (const char of line) {
            if (char === '"') {
                inQuotes = !inQuotes;
            } else if (char === ',' && !inQuotes) {
                values.push(current.trim());
                current = '';
            } else {
                current += char;
            }
        }
        values.push(current.trim());
        
        rows.push(values);
    }
    
    return [headers, ...rows];
}

// Calculate age as of July 1, 2025
function calculateAge(birthDateStr: string): number {
    const targetDate = new Date('2025-07-01');
    const birthDate = new Date(birthDateStr);
    
    let age = targetDate.getFullYear() - birthDate.getFullYear();
    
    const monthDiff = targetDate.getMonth() - birthDate.getMonth();
    const dayDiff = targetDate.getDate() - birthDate.getDate();
    
    if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
        age--;
    }
    
    return age;
}

// Main transformation
function main() {
    const csvPath = path.join('input', 'input.csv');
    const expectedPath = path.join('input', 'expected_format.json');
    
    // Read CSV
    const csvData = readCSV(csvPath);
    const headers = csvData[0];
    const rows = csvData.slice(1);
    
    // Read expected format to understand structure
    const expectedContent = fs.readFileSync(expectedPath, 'utf-8');
    const expectedFormat = JSON.parse(expectedContent);
    
    // Build output based on expected format structure
    const output: any[] = [];
    
    for (const row of rows) {
        const record: any = {};
        
        for (let i = 0; i < headers.length; i++) {
            const header = headers[i];
            const value = row[i] || '';
            
            // Map headers based on expected format
            if (header === 'id') {
                record.id = parseInt(value, 10);
            } else if (header === 'name') {
                record.name = value;
            } else if (header === 'email') {
                record.email = value;
            } else if (header === 'birth_date' || header === 'dob') {
                record.age = calculateAge(value);
            } else if (header === 'department' || header === 'dept') {
                record.department = value;
            } else if (header === 'salary') {
                record.salary = parseFloat(value);
            } else if (header === 'active' || header === 'status') {
                record.isActive = value.toLowerCase() === 'true' || value === '1' || value === 'yes';
            } else {
                // Default mapping based on expected format
                const mappedKey = header.toLowerCase().replace(/_/g, '');
                record[mappedKey] = isNaN(parseFloat(value)) ? value : parseFloat(value);
            }
        }
        
        output.push(record);
    }
    
    // Output JSON array
    console.log(JSON.stringify(output, null, 2));
}

main();