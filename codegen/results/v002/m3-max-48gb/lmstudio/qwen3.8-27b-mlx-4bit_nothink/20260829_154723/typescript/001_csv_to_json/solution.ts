import * as fs from 'fs';

function parseCsvLine(line: string): string[] {
  const result: string[] = [];
  let current = '';
  let inQuotes = false;
  for (let i = 0; i < line.length; i++) {
    const c = line[i];
    if (inQuotes) {
      if (c === '"') {
        if (i + 1 < line.length && line[i + 1] === '"') {
          current += '"';
          i++;
        } else {
          inQuotes = false;
        }
      } else {
        current += c;
      }
    } else {
      if (c === '"') {
        inQuotes = true;
      } else if (c === ',') {
        result.push(current);
        current = '';
      } else {
        current += c;
      }
    }
  }
  result.push(current);
  return result;
}

function main() {
  const content = fs.readFileSync('input/data.csv', 'utf8');
  const lines = content.split(/\r?\n/).filter(l => l.length > 0);
  if (lines.length === 0) {
    process.stdout.write('[]');
    return;
  }
  const headers = parseCsvLine(lines[0]).map(h => h.trim());
  const rows = lines.slice(1);
  const objects: any[] = [];
  for (const row of rows) {
    const cols = parseCsvLine(row).map(c => c.trim());
    const obj: any = {};
    for (let i = 0; i < headers.length; i++) {
      obj[headers[i]] = cols[i] !== undefined ? cols[i] : '';
    }
    const record = {
      Name: String(obj['Name'] ?? ''),
      Age: parseInt(String(obj['Age'] ?? '0'), 10),
      Email: String(obj['Email'] ?? ''),
      Score: parseFloat(String(obj['Score'] ?? '0')),
    };
    objects.push(record);
  }
  process.stdout.write(JSON.stringify(objects));
}

main();