import * as fs from 'fs';

function main() {
  const content = fs.readFileSync('input/data.csv', 'utf-8');
  const lines = content.split(/\r?\n/).filter(line => line.trim().length > 0);
  
  const result = lines.slice(1).map(line => {
    const [name, ageStr, email, scoreStr] = line.split(',').map(val => val.trim());
    return {
      Name: name,
      Age: parseInt(ageStr, 10),
      Email: email,
      Score: parseFloat(scoreStr)
    };
  });
  
  process.stdout.write(JSON.stringify(result));
  process.stdout.write('\n');
}

main();