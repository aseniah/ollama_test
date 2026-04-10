import * as fs from 'fs';
import * as path from 'path';

const inputFile = 'input/data.csv';

if (process.argv[2] && process.argv[2] !== inputFilename) {
    // Use CLI argument if provided
    const inputFileArg = process.argv[2];
    const realInput = path.resolve(inputFileArg);
} else {
    const realInput = path.resolve(inputFile);
}

const fileContent = fs.readFileSync(realInput, 'utf-8');
const lines = fileContent.trim().split('\n');
const [headerLine, ...dataLines] = lines;

const [header] = headerLine.split(',').map((v) => v.trim());

const result: Array<{
    name: string;
    age: number;
    email: string;
    score: number;
}> = [];

for (const line of dataLines) {
    if (!line.trim()) continue;
    const [name, ageStr, email, scoreStr] = line.split(',');
    result.push({
        name: name.trim(),
        age: parseInt(ageStr.trim(), 10),
        email: email.trim(),
        score: parseFloat(scoreStr.trim()),
    });
}

console.log(JSON.stringify(result, null, 2));