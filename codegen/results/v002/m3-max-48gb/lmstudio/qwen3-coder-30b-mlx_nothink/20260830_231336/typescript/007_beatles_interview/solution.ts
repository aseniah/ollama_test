import { readFileSync } from 'fs';
import { join } from 'path';

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

const calculateAge = (birthday: string, referenceDate: Date): number => {
  const birthDate = new Date(birthday);
  const age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    return age - 1;
  }
  
  return age;
};

const parseDate = (dateStr: string): string => {
  const parts = dateStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
};

const processCSV = (): Person[] => {
  const csvContent = readFileSync(join('input', 'input.csv'), 'utf-8');
  const lines = csvContent.split('\n').slice(1); // Skip header
  
  const people: Person[] = [];
  const referenceDate = new Date('2025-07-01');
  
  for (const line of lines) {
    if (!line.trim()) continue;
    
    const [name, birthday, died, father, mother, brother, sister] = line.split(',');
    
    // Parse name
    const nameParts = name.trim().split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts.slice(1).join(' ');
    
    // Parse birthday
    const parsedBirthday = parseDate(birthday.trim());
    
    // Calculate age
    const age = calculateAge(parsedBirthday, referenceDate);
    
    // Build relatives
    const relatives: Relative[] = [];
    
    if (father.trim() !== 'null') {
      const [fFirstName, fLastName] = father.trim().split(' ');
      relatives.push({
        FirstName: fFirstName,
        LastName: fLastName,
        Relationship: 'Father'
      });
    }
    
    if (mother.trim() !== 'null') {
      const [mFirstName, mLastName] = mother.trim().split(' ');
      relatives.push({
        FirstName: mFirstName,
        LastName: mLastName,
        Relationship: 'Mother'
      });
    }
    
    if (brother.trim() !== 'null') {
      const [bFirstName, bLastName] = brother.trim().split(' ');
      relatives.push({
        FirstName: bFirstName,
        LastName: bLastName,
        Relationship: 'Brother'
      });
    }
    
    if (sister.trim() !== 'null') {
      const [sFirstName, sLastName] = sister.trim().split(' ');
      relatives.push({
        FirstName: sFirstName,
        LastName: sLastName,
        Relationship: 'Sister'
      });
    }
    
    people.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: parsedBirthday,
      Age: age,
      Relatives: relatives
    });
  }
  
  return people;
};

const people = processCSV();
console.log(JSON.stringify(people, null, 2));