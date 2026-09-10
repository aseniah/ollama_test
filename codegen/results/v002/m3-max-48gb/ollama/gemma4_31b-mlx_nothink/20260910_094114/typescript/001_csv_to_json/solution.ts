import * as fs from 'fs';
import * as path from 'path';

interface UserData {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

function main() {
    const filePath = path.join('input', 'data.csv');
    
    try {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        const lines = fileContent.split(/\r?\n/).filter(line => line.trim() !== '');
        
        if (lines.length === 0) {
            console.log(JSON.stringify([]));
            return;
        }

        // Extract headers (Name,Age,Email,Score)
        const headers = lines[0].split(',');
        const dataRows = lines.slice(1);

        const result: UserData[] = dataRows.map(line => {
            const values = line.split(',');
            
            return {
                Name: values[0] ? values[0].trim() : '',
                Age: values[1] ? parseInt(values[1].trim(), 10) : 0,
                Email: values[2] ? values[2].trim() : '',
                Score: values[3] ? parseFloat(values[3].trim()) : 0.0
            };
        });

        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
    } catch (error) {
        // Silent failure or handle as needed per requirements
        process.exit(1);
    }
}

main();