import * as fs from 'fs';
import * as path from 'path';

function parseCSV(content: string): string[][] {
  const lines = content.split('\n').filter(line => line.trim() !== '');
  return lines.map(line => line.split(','));
}

function calculateAge(birthday: string, referenceDate: Date): number {
  // birthday is in MM/DD/YYYY format
  const [month, day, year] = birthday.split('/').map(Number);
  const birth = new Date(year, month - 1, day);
  
  let age = referenceDate.getFullYear() - birth.getFullYear();
  const monthDiff = referenceDate.getMonth() - birth.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  const rows = parseCSV(content);
  const header = rows[0];
  const dataRows = rows.slice(1);
  
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed)
  
  const result: any[] = [];
  
  for (const row of dataRows) {
    const name = row[0];
    const birthday = row[1];
    const died = row[2];
    const father = row[3];
    const mother = row[4];
    const brother = row[5];
    const sister = row[6];
    
    // Split name into FirstName and LastName
    const nameParts = name.split(' ');
    const lastName = nameParts.pop() || '';
    const firstName = nameParts.join(' ');
    
    const age = calculateAge(birthday, referenceDate);
    
    const relatives: any[] = [];
    
    if (father && father !== 'null') {
      const fatherParts = father.split(' ');
      relatives.push({
        FirstName: fatherParts[0],
        LastName: fatherParts.slice(1).join(' '),
        Relationship: 'Father'
      });
    }
    
    if (mother && mother !== 'null') {
      const motherParts = mother.split(' ');
      relatives.push({
        FirstName: motherParts[0],
        LastName: motherParts.slice(1).join(' '),
        Relationship: 'Mother'
      });
    }
    
    if (brother && brother !== 'null') {
      const brotherParts = brother.split(' ');
      relatives.push({
        FirstName: brotherParts[0],
        LastName: brotherParts.slice(1).join(' '),
        Relationship: 'Brother'
      });
    }
    
    if (sister && sister !== 'null') {
      const sisterParts = sister.split(' ');
      relatives.push({
        FirstName: sisterParts[0],
        LastName: sisterParts.slice(1).join(' '),
        Relationship: 'Sister'
      });
    }
    
    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthday.split('/').reverse().join('-'),
      Age: age,
      Relatives: relatives
    });
  }
  
  console.log(JSON.stringify(result, null, 2));
}

main();