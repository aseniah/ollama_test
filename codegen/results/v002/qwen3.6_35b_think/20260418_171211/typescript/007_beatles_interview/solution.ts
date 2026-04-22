import * as fs from 'fs';
import * as path from path;

const csv = fs.readFileSync(path.join(__dirname, 'input', 'input.csv'), 'utf-8');
const lines = csv.trim().split('\n');
const header = lines[0].split(',');

const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed in JS)

const result: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const cols = lines[i].split(',');
  const name = cols[0].trim();
  const birthdayStr = cols[1].trim();
  const died = cols[2].trim();
  const father = cols[3].trim();
  const mother = cols[4].trim();
  const brother = cols[5].trim();
  const sister = cols[6].trim();

  const parts = name.split(' ');
  const firstName = parts[0];
  const lastName = parts.slice(1).join(' ');

  const [birthMonth, birthDay, birthYear] = birthdayStr.split('/').map(Number);
  const birthday = new Date(birthYear, birthMonth - 1, birthDay);

  const birthdayFormatted = `${birthYear}-${String(birthMonth).padStart(2, '0')}-${String(birthDay).padStart(2, '0')}`;

  let age = referenceDate.getFullYear() - birthday.getFullYear();
  if (
    referenceDate.getMonth() < birthday.getMonth() ||
    (referenceDate.getMonth() === birthday.getMonth() && referenceDate.getDate() < birthday.getDate())
  ) {
    age--;
  }

  const relatives: any[] = [];

  if (father !== 'null') {
    const fParts = father.split(' ');
    relatives.push({
      FirstName: fParts[0],
      LastName: fParts.slice(1).join(' '),
      Relationship: 'Father'
    });
  }
  if (mother !== 'null') {
    const mParts = mother.split(' ');
    relatives.push({
      FirstName: mParts[0],
      LastName: mParts.slice(1).join(' '),
      Relationship: 'Mother'
    });
  }
  if (brother !== 'null') {
    const bParts = brother.split(' ');
    relatives.push({
      FirstName: bParts[0],
      LastName: bParts.slice(1).join(' '),
      Relationship: 'Brother'
    });
  }
  if (sister !== 'null') {
    const sParts = sister.split(' ');
    relatives.push({
      FirstName: sParts[0],
      LastName: sParts.slice(1).join(' '),
      Relationship: 'Sister'
    });
  }

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayFormatted,
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));