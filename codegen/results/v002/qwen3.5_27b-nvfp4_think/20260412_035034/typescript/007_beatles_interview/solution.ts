import * as fs from 'fs';
import * as path from 'path';

const REFERENCE_DATE = new Date('2025-07-01');

function parseBirthday(dateStr: string): string {
  const parts = dateStr.split('/').map(Number);
  const month = String(parts[0]).padStart(2, '0');
  const day = String(parts[1]).padStart(2, '0');
  const year = String(parts[2]);
  return `${year}-${month}-${day}`;
}

function calculateAge(birthdayStr: string, deathDate?: Date): number {
  const birth = new Date(birthdayStr);
  let targetDate = REFERENCE_DATE;
  
  if (deathDate) {
    targetDate = deathDate!;
  }
  
  let age = targetDate.getFullYear() - birth.getFullYear();
  const monthDiff = targetDate.getMonth() - birth.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && targetDate.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  const firstName = parts[0];
  const lastName = parts[parts.length - 1];
  return { firstName, lastName };
}

function parseRelatives(personName: string): Array<{ FirstName: string; LastName: string; Relationship: string }> {
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  
  // Father and Mother should always be included (based on expected output)
  // Check the CSV columns and add them to relatives array
  
  return relatives;
}

function main(): void {
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  const lines = content.trim().split('\n');
  const headers = lines[0].split(',').map(h => h.trim());
  const result: any[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',').map(v => v.trim());
    
    // Map headers to indices
    const headerIndex: Record<string, number> = {};
    headers.forEach((h, idx) => {
      headerIndex[h] = idx;
    });
    
    const fullName = values[headerIndex['Name']];
    const birthdayStr = values[headerIndex['Birthday']];
    const diedStr = values[headerIndex['Died']];
    const fatherStr = values[headerIndex['Father']];
    const motherStr = values[headerIndex['Mother']];
    const brotherStr = values[headerIndex['Brother']];
    const sisterStr = values[headerIndex['Sister']];
    
    const { firstName, lastName } = parseName(fullName);
    const formattedBirthday = parseBirthday(birthdayStr);
    
    let deathDate: Date | undefined;
    if (diedStr && diedStr !== 'null') {
      deathDate = new Date(diedStr.replace(/(\d+)\/(\d+)\/(\d+)/, '$3-$1-$2'));
    }
    
    const age = calculateAge(formattedBirthday, deathDate);
    
    const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
    
    if (fatherStr && fatherStr !== 'null') {
      const rel = parseName(fatherStr);
      relatives.push({ FirstName: rel.firstName, LastName: rel.lastName, Relationship: 'Father' });
    }
    
    if (motherStr && motherStr !== 'null') {
      const rel = parseName(motherStr);
      relatives.push({ FirstName: rel.firstName, LastName: rel.lastName, Relationship: 'Mother' });
    }
    
    if (brotherStr && brotherStr !== 'null') {
      const rel = parseName(brotherStr);
      relatives.push({ FirstName: rel.firstName, LastName: rel.lastName, Relationship: 'Brother' });
    }
    
    if (sisterStr && sisterStr !== 'null') {
      const rel = parseName(sisterStr);
      relatives.push({ FirstName: rel.firstName, LastName: rel.lastName, Relationship: 'Sister' });
    }
    
    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    });
  }
  
  console.log(JSON.stringify(result, null, '  '));
}

main();