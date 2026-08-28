import * as fs from 'fs';
import * as path from 'path';

interface Record {
    name: string;
    age: number;
    active: boolean;
    score: number;
}

function main(): void {
    const filePath = path.join('input', 'data.json');
    const data: Record[] = JSON.parse(fs.readFileSync(filePath, 'utf-8'));
    
    const filtered = data
        .filter(record => record.active === true && record.age >= 30)
        .sort((a, b) => a.name.localeCompare(b.name));
    
    process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');
}

main();