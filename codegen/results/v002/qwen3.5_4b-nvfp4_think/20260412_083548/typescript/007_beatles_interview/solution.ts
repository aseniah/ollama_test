import * as fs from 'fs';

const referenceDate = new Date('2025-07-01');

function parseBirthday(str: string): Date {
    const [month, day, year] = str.split('/').map(Number);
    return new Date(`${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`);
}

function calculateAge(birthdayStr: string): number {
    const birthday = parseBirthday(birthdayStr);
    const ageAsOfJuly1 = new Date('2025-07-01');
    
    let years = ageAsOfJuly1.getFullYear() - birthday.getFullYear();
    const birthdayMonthDay = new Date(`${year}=${String(year).padStart(2, '0')}-${String(day).padStart(2, '0')}`).toISOString().split('T')[0];
    
    return years;
}