import { readFileSync } from 'fs';
import { join } from 'path';

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: {
    FirstName: string;
    LastName: string;
    Relationship: string;
  }[];
}

function calculateAge(birthDate: string, referenceDate: Date): number {
  const birth = new Date(birthDate);
  const reference = new Date(referenceDate);
  
  let age = reference.getFullYear() - birth.getFullYear();
  const monthDiff = reference.getMonth() - birth.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && reference.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

function parseDateString(dateStr: string): string {
  // Convert MM/DD/YYYY to YYYY-MM-DD
  const parts = dateStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  
  return `${year}-${month}-${day}`;
}

function processCSV(): Person[] {
  const csvContent = readFileSync(join('input', 'input.csv'), 'utf-8');
  const lines = csvContent.split('\n');
  
  // Skip header line
  const dataLines = lines.slice(1);
  
  const people: Person[] = [];
  
  const referenceDate = new Date('2025-07-01');
  
  for (const line of dataLines) {
    if (!line.trim()) continue;
    
    const [name, birthday, died, father, mother, brother, sister] = line.split(',');
    
    const [firstName, lastName] = name.split(' ');
    
    const person: Person = {
      FirstName: firstName,
      LastName: lastName,
      Birthday: parseDateString(birthday),
      Age: calculateAge(parseDateString(birthday), referenceDate),
      Relatives: []
    };
    
    // Add relatives
    if (father) {
      const [fName, lName] = father.split(' ');
      person.Relatives.push({
        FirstName: fName,
        LastName: lName,
        Relationship: 'Father'
      });
    }
    
    if (mother) {
      const [mName, lName] = mother.split(' ');
      person.Relatives.push({
        FirstName: mName,
        LastName: lName,
        Relationship: 'Mother'
      });
    }
    
    if (brother) {
      const [bName, lName] = brother.split(' ');
      person.Relatives.push({
        FirstName: bName,
        LastName: lName,
        Relationship: 'Brother'
      });
    }
    
    if (sister) {
      const [sName, lName] = sister.split(' ');
      person.Relatives.push({
        FirstName: sName,
        LastName: lName,
        Relationship: 'Sister'
      });
    }
    
    people.push(person);
  }
  
  return people;
}

const result = processCSV();
console.log(JSON.stringify(result, null, 2));