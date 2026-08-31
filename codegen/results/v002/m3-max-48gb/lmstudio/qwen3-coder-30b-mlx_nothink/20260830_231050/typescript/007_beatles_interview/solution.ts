import fs from 'fs';
import path from 'path';

// Helper function to calculate age as of July 1, 2025
function calculateAge(birthDate: Date): number {
  const referenceDate = new Date('2025-07-01');
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  
  // Adjust if birthday hasn't occurred yet this year
  if (referenceDate.getMonth() < birthDate.getMonth() || 
      (referenceDate.getMonth() === birthDate.getMonth() && 
       referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Helper function to parse date from MM/DD/YYYY format
function parseDate(dateStr: string): Date {
  const [month, day, year] = dateStr.split('/').map(Number);
  return new Date(year, month - 1, day);
}

// Helper function to parse name
function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.split(' ');
  const firstName = parts[0];
  const lastName = parts.slice(1).join(' ');
  return { firstName, lastName };
}

// Read and parse CSV file
const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const lines = csvContent.split('\n').slice(1); // Skip header

const result = lines
  .filter(line => line.trim() !== '')
  .map(line => {
    const [name, birthday, died, father, mother, brother, sister] = line.split(',');
    
    const { firstName, lastName } = parseName(name.trim());
    const birthDate = parseDate(birthday.trim());
    
    const age = calculateAge(birthDate);
    
    const relatives: any[] = [];
    
    if (father.trim() !== 'null') {
      const { firstName: fatherFirstName, lastName: fatherLastName } = parseName(father.trim());
      relatives.push({
        FirstName: fatherFirstName,
        LastName: fatherLastName,
        Relationship: "Father"
      });
    }
    
    if (mother.trim() !== 'null') {
      const { firstName: motherFirstName, lastName: motherLastName } = parseName(mother.trim());
      relatives.push({
        FirstName: motherFirstName,
        LastName: motherLastName,
        Relationship: "Mother"
      });
    }
    
    if (brother.trim() !== 'null') {
      const { firstName: brotherFirstName, lastName: brotherLastName } = parseName(brother.trim());
      relatives.push({
        FirstName: brotherFirstName,
        LastName: brotherLastName,
        Relationship: "Brother"
      });
    }
    
    if (sister.trim() !== 'null') {
      const { firstName: sisterFirstName, lastName: sisterLastName } = parseName(sister.trim());
      relatives.push({
        FirstName: sisterFirstName,
        LastName: sisterLastName,
        Relationship: "Sister"
      });
    }
    
    return {
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthDate.toISOString().split('T')[0],
      Age: age,
      Relatives: relatives
    };
  });

console.log(JSON.stringify(result, null, 2));