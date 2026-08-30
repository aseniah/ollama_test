import fs from 'fs';
import path from 'path';

const inputFilePath = path.resolve(__dirname, 'input/input.csv');
const outputFilePath = path.resolve(__dirname, 'output.json');

async function readInputAndCalculateAges() {
  const inputContent = await fs.readFileSync(inputFilePath, 'utf8');
  const rows = inputContent.trim().split('\n').filter(row => row.trim() !== '');

  const names = new Map<string, string>();
  const birthDatums = new Map<string, Date>();
  const relatives = new Map<string, string>();
  const fathers = new Map<string, string>();
  const mothers = new Map<string, string>();
  const brothers = new Map<string, string>();
  const sisters = new Map<string, string>();

  // Parse input data
  for (const row of rows) {
    const columns = row.split(',');
    if (columns.length < 8) continue; // Skip incomplete rows

    const name = columns[0].trim();
    const birthday = columns[1].trim();
    const died = columns[2].trim();
    const father = columns[3].trim();
    const mother = columns[4].trim();
    const brother = columns[5].trim();
    const sister = columns[6].trim();

    names.set(name, name);
    birthDatums.set(name, parseDate(birthday));
    relatives.set(name, father ? father : 'null');
    fathers.set(name, father ? father : 'null');
    mothers.set(name, mother ? mother : 'null');
    brothers.set(name, brother ? brother : 'null');
    sisters.set(name, sister ? sister : 'null');
  }

  // Calculate ages
  const result: Array<Record<string, any>> = [];
  for (const [name, birthDate] of birthDatums) {
    const currentDate = new Date(2025, 6, 1); // July 1, 2025
    const age = currentDate.getFullYear() - birthDate.getFullYear() - ((currentDate.getMonth() + 1) < birthDate.getMonth() ? 1 : 0);
    const relativesInfo = relatives.get(name) || 'null';
    const fathersInfo = fathers.get(name) || 'null';
    const mothersInfo = mothers.get(name) || 'null';
    const brothersInfo = brothers.get(name) || 'null';
    const sistersInfo = sisters.get(name) || 'null';

    result.push({
      FirstName: name,
      LastName: names.get(name),
      Birthday: birthDate.toISOString(),
      Age: age,
      Relatives: {
        ...relativesInfo.split(',').map(rel => ({
          FirstName: rel.split(' ')[0].trim(),
          LastName: rel.split(' ')[1].trim(),
          Relationship: relativesInfo.includes(rel) ? 'null' : relativesInfo.includes(rel) ? relativesInfo.split(',')[rel.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim() : relativesInfo.split(',')[rel.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim(),
          Relationship: relativesInfo.includes(rel) ? relativesInfo.split(',')[rel.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim() : relativesInfo.split(',')[rel.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim()
        })),
        ...fathersInfo.split(',')
          .map(father => ({
            FirstName: father.split(' ')[0].trim(),
            LastName: father.split(' ')[1].trim(),
            Relationship: fathersInfo.includes(father) ? 'null' : fathersInfo.includes(father) ? fathersInfo.split(',')[father.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim() : fathersInfo.split(',')[father.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim(),
            Relationship: fathersInfo.includes(father) ? fathersInfo.split(',')[father.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim() : fathersInfo.split(',')[father.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim()
          })),
        ...mothersInfo.split(',')
          .map(mother => ({
            FirstName: mother.split(' ')[0].trim(),
            LastName: mother.split(' ')[1].trim(),
            Relationship: mothersInfo.includes(mother) ? 'null' : mothersInfo.includes(mother) ? mothersInfo.split(',')[mother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim() : mothersInfo.split(',')[mother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim(),
            Relationship: mothersInfo.includes(mother) ? mothersInfo.split(',')[mother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim() : mothersInfo.split(',')[mother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim()
          })),
        ...brothersInfo.split(',')
          .map(brother => ({
            FirstName: brother.split(' ')[0].trim(),
            LastName: brother.split(' ')[1].trim(),
            Relationship: brothersInfo.includes(brother) ? 'null' : brothersInfo.includes(brother) ? brothersInfo.split(',')[brother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim() : brothersInfo.split(',')[brother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim(),
            Relationship: brothersInfo.includes(brother) ? brothersInfo.split(',')[brother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim() : brothersInfo.split(',')[brother.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim()
          })),
        ...sistersInfo.split(',')
          .map(sister => ({
            FirstName: sister.split(' ')[0].trim(),
            LastName: sister.split(' ')[1].trim(),
            Relationship: sistersInfo.includes(sister) ? 'null' : sistersInfo.includes(sister) ? sistersInfo.split(',')[sister.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim() : sistersInfo.split(',')[sister.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[0].trim(),
            Relationship: sistersInfo.includes(sister) ? sistersInfo.split(',')[sister.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim() : sistersInfo.split(',')[sister.split(',')[0].split(' ').indexOf('(') - 1].split(' ')[1].trim()
          })),
      },
    });
  }

  fs.writeFileSync(outputFilePath, JSON.stringify(result, null, 2));
}

async function parseDate(dateString: string): Date {
  const parts = dateString.split('-');
  if (parts.length !== 3) throw new Error('Invalid date format');

  const year = parseInt(parts[2], 10);
  const month = parseInt(parts[1], 10) - 1; // Convert to 0-based index
  const day = parseInt(parts[0], 10);

  const date = new Date(year, month, day);
  return date;
}

readInputAndCalculateAges();