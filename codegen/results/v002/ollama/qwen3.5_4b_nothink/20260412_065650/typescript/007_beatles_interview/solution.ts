import { readFile } from 'fs/promises';
import { join } from 'path';

const referenceDate = new Date('2025-07-01');

async function main() {
  const inputPath = join(process.argv[2], 'input.csv');
  const outputPath = join(process.argv[2], 'input/expected_format.json');

  const csvContent = await readFile(inputPath, 'utf-8');
  const lines = csvContent.trim().split('\n');
  const headers = lines[0].split(',').map((h, i) => i); // ['Name', 'Birthday', 'Died', 'Father', 'Mother', 'Brother', 'Sister']

  const result: any[] = [];

  for (let i = 1; i < lines.length; i++) {
    const row = lines[i].split(',');
    const name = row[headers[0]].trim();
    const birthdayStr = row[headers[1]].trim();
    
    // Parse birthday: MM/DD/YYYY
    const [birthMonth, birthDay, birthYear] = birthdayStr.split('/').map(Number);
    const birthDate = new Date(birthYear, birthMonth - 1, birthDay);

    // Extract relatives from columns
    const relatives = [];

    // Father (column index 3)
    if (row[3].trim() !== 'null') {
      relatives.push({
        FirstName: row[3].trim().split(' ')[0],
        LastName: row[3].trim().split(' ').pop(),
        Relationship: 'Father'
      });
    }

    // Mother (column index 4)
    if (row[4].trim() !== 'null') {
      relatives.push({
        FirstName: row[4].trim().split(' ')[0],
        LastName: row[4].trim().split(' ').pop(),
        Relationship: 'Mother'
      });
    }

    // Brother (column index 5)
    if (row[5].trim() !== 'null') {
      relatives.push({
        FirstName: row[5].trim().split(' ')[0],
        LastName: row[5].trim().split(' ').pop(),
        Relationship: 'Brother'
      });
    }

    // Sister (column index 6)
    if (row[6].trim() !== 'null') {
      relatives.push({
        FirstName: row[6].trim().split(' ')[0],
        LastName: row[6].trim().split(' ').pop(),
        Relationship: 'Sister'
      });
    }

    // Calculate age as of July 1, 2025
    let age = referenceDate.getFullYear() - birthDate.getFullYear();
    const refMonth = 6, refDay = 1;
    const birthMonth = birthDate.getMonth() + 1, birthDay = birthDate.getDate();

    if (refMonth < birthMonth || (refMonth === birthMonth && refDay < birthDay)) {
      age -= 1;
    }

    result.push({
      FirstName: name.split(' ').pop(), // Last name for Beatles members (they share same last name usually)
      LastName: name.split(' ').pop(),
      Birthday: `${String(birthMonth).padStart(2, '0')}/${String(birthDay).padStart(2, '0')}/${birthYear}`,
      Age: age,
      Relatives: relatives
    });
  }

  // Use simple JSON.stringify as we don't have a JSON module
  console.log(JSON.stringify(result, null, 2));
}

main();