import * as fs from 'fs';
import * as path from 'path';

interface DataRecord {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

function main(): void {
    const inputPath = path.resolve('input/data.csv');
    
    let content: string;
    try {
        content = fs.readFileSync(inputPath, 'utf-8');
    } catch (error) {
        console.error(`Error reading file: ${error}`);
        process.exit(1);
    }

    // Split by newlines, trim each line
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    if (lines.length === 0) {
        console.log('[]');
        return;
    }

    // First line is header, skip it
    const dataLines = lines.slice(1);
    const records: DataRecord[] = [];

    for (const line of dataLines) {
        const fields = line.split(',');
        if (fields.length < 4) continue;

        const name = fields[0].trim();
        const age = parseInt(fields[1].trim(), 10);
        const email = fields[2].trim();
        const score = parseFloat(fields[3].trim());

        records.push({
            Name: name,
            Age: age,
            Email: email,
            Score: score
        });
    }

    console.log(JSON.stringify(records, null, 2));
}

main();