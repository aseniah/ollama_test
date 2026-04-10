import { readFileSync } from 'fs';
import { resolve } from 'path';

const filePath = resolve(process.cwd(), 'input/data.csv');

try {
    const fileContent = readFileSync(filePath, 'utf-8');
    const lines = fileContent.split('\n');
    const header = lines[0].split(',');
    const data: any[] = [];

    for (let i = 1; i < lines.length; i++) {
        const row = lines[i].split(',');
        if (row.length !== header.length) continue;
        const obj: any = {};
        for (let j = 0; j < header.length; j++) {
            switch (header[j]) {
                case 'Name':
                    obj.Name = row[j];
                    break;
                case 'Age':
                    obj.Age = parseInt(row[j], 10);
                    break;
                case 'Email':
                    obj.Email = row[j];
                    break;
                case 'Score':
                    obj.Score = parseFloat(row[j]);
                    break;
            }
        }
        data.push(obj);
    }

    console.log(JSON.stringify(data, null, 2));
} catch (error) {
    console.error('Error reading file:', error.message);
}