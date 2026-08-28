import * as fs from 'fs';
import * as path from 'path';

function calculateAge(birthdayStr: string, referenceDate: Date): number {
  const parts = birthdayStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);

  const birthday = new Date(year, month - 1, day);
  
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthday.getMonth();
  const dayDiff = referenceDate.getDate() - birthday.getDate();
  
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age -= 1;
  }
  
  return age;
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  const lines = content.trim().split('\n');
  const headers = lines[0].split(',');
  
  const referenceDate = new Date(Date.UTC(2025, 6, 1)); // July 1, 2025
  
  const result: any[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i];
    // Split by comma, but need to handle potential issues. Assuming no commas in fields based on sample.
    const values = line.split(',');
    
    const nameStr = values[0].trim();
    const birthdayStr = values[1].trim();
    // const died = values[2].trim(); // Not used
    
    const fatherStr = values[3].trim();
    const motherStr = values[4].trim();
    const brotherStr = values[5].trim();
    const sisterStr = values[6].trim();
    
    // Parse name: "FirstName LastName"
    // Some names have middle names like "John Winston Lennon". 
    // The expected output shows FirstName: "John", LastName: "Lennon".
    // So first word is FirstName, last word is LastName.
    const nameParts = nameStr.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];
    
    const age = calculateAge(birthdayStr, referenceDate);
    
    const relatives: any[] = [];
    
    if (fatherStr && fatherStr !== 'null') {
      const fParts = fatherStr.split(' ');
      relatives.push({
        FirstName: fParts[0],
        LastName: fParts[fParts.length - 1],
        Relationship: "Father"
      });
    }
    
    if (motherStr && motherStr !== 'null') {
      const mParts = motherStr.split(' ');
      relatives.push({
        FirstName: mParts[0],
        LastName: mParts[mParts.length - 1],
        Relationship: "Mother"
      });
    }
    
    if (brotherStr && brotherStr !== 'null') {
      const bParts = brotherStr.split(' ');
      relatives.push({
        FirstName: bParts[0],
        LastName: bParts[bParts.length - 1],
        Relationship: "Brother"
      });
    }
    
    if (sisterStr && sisterStr !== 'null') {
      const sParts = sisterStr.split(' ');
      relatives.push({
        FirstName: sParts[0],
        LastName: sParts[sParts.length - 1],
        Relationship: "Sister"
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
  
  console.log(JSON.stringify(result, null, 2));
}

main();