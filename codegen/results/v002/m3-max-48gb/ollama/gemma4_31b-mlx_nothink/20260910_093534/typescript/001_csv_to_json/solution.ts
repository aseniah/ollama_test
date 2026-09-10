import * as fs from 'fs';
import * as path from 'path';

interface UserData {
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
        
        if (lines.length === 0) {
            console.log(JSON.stringify([]));
            return;
        }

        // The first line is the header, we skip it
        const rows = lines.slice(1);
        const result: UserData[] = [];

        for (const row of rows) {
            const columns = row.split(',');
            if (columns.length >= 4) {
                result.push({
                    Name: columns[0].trim(),
                    Age: parseInt(columns[1].trim(), 10),
                    Email: columns[2].trim(),
                    Score: parseFloat(columns[3].trim())
                });
            }
        }

        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
    } catch (error) {
        // Handle file not found or other read errors silently or as an empty array
        process.stdout.write(JSON.stringify([]) + '\n');
    }
}

solve();