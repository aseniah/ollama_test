import * as fs from 'fs';
import * as path from 'path';

interface User {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

function solve() {
    const filePath = path.join('input', 'data.csv');
    
    try {
        const content = fs.readFileSync(filePath, 'utf8');
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
        
        // Remove header row
        const dataRows = lines.slice(1);
        
        const result: User[] = dataRows.map(row => {
            const columns = row.split(',');
            
            return {
                Name: columns[0].trim(),
                Age: parseInt(columns[1].trim(), 10),
                Email: columns[2].trim(),
                Score: parseFloat(columns[3].trim())
            };
        });

        process.stdout.write(JSON.stringify(result, null, 2));
    } catch (error) {
        // Silent failure as per requirement to only output valid JSON 
        // or the resulting array.
    }
}

solve();