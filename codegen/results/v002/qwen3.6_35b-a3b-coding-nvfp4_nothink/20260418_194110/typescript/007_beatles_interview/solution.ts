import * as fs from 'fs';
import * as path from 'path';

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

function parseDate(dateStr: string): Date {
  const parts = dateStr.split('/');
  if (parts.length !== 3) {
    throw new Error(`Invalid date format: ${dateStr}`);
  }
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

function calculateAge(birthday: Date, asOf: Date): number {
  let age = asOf.getFullYear() - birthday.getFullYear();
  const monthDiff = asOf.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && asOf.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function main() {
  const inputPath = path.join(__dirname, 'input', 'input.csv');
  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  
  const lines = csvContent.trim().split('\n');
  const header = lines[0].split(',');
  
  const asOfDate = new Date(2025, 6, 1); // July 1, 2025
  
  const results: Person[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;
    
    // Split by comma, but be careful with potential commas in names? 
    // The data seems simple enough for comma split.
    const fields = line.split(',');
    
    const nameField = fields[0].trim();
    const birthdayStr = fields[1].trim();
    const diedStr = fields[2].trim();
    const fatherStr = fields[3].trim();
    const motherStr = fields[4].trim();
    const brotherStr = fields[5].trim();
    const sisterStr = fields[6].trim();
    
    const nameParts = nameField.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts.slice(1).join(' ');
    
    const birthdayDate = parseDate(birthdayStr);
    const age = calculateAge(birthdayDate, asOfDate);
    
    const relatives: Relative[] = [];
    
    if (fatherStr !== 'null' && fatherStr !== '') {
      const fParts = fatherStr.split(' ');
      relatives.push({
        FirstName: fParts[0],
        LastName: fParts.slice(1).join(' '),
        Relationship: 'Father'
      });
    }
    
    if (motherStr !== 'null' && motherStr !== '') {
      const mParts = motherStr.split(' ');
      relatives.push({
        FirstName: mParts[0],
        LastName: mParts.slice(1).join(' '),
        Relationship: 'Mother'
      });
    }
    
    if (brotherStr !== 'null' && brotherStr !== '') {
      const bParts = brotherStr.split(' ');
      relatives.push({
        FirstName: bParts[0],
        LastName: bParts.slice(1).join(' '),
        Relationship: 'Brother'
      });
    }
    
    if (sisterStr !== 'null' && sisterStr !== '') {
      const sParts = sisterStr.split(' ');
      relatives.push({
        FirstName: sParts[0],
        LastName: sParts.slice(1).join(' '),
        Relationship: 'Sister'
      });
    }
    
    results.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayStr,
      Age: age,
      Relatives: relatives
    });
  }
  
  console.log(JSON.stringify(results, null, 2));
}

main();