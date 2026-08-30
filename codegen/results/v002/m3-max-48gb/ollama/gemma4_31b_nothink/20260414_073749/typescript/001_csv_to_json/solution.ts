import * as fs from 'fs';
import * as path from 'path';

interface UserData {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  try {
    const filePath = path.join('input', 'data.csv');
    const content = fs.readFileSync(filePath, 'utf-8');
    
    // Split file into lines and filter out empty lines
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    // Remove header row
    const dataLines = lines.slice(1);
    
    const result: UserData[] = dataLines.map(line => {
      const [name, age, email, score] = line.split(',');
      
      return {
        Name: name,
        Age: parseInt(age, 10),
        Email: email,
        Score: parseFloat(score),
      };
    });

    process.stdout.write(JSON.stringify(result, null, 2));
  } catch (error) {
    // Silently exit or handle error according to requirements
    process.exit(1);
  }
}

main();