import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header
const header = lines[0];
const rows: Array<Record<string, string>> = [];

for (let i = 1; i < lines.length; i++) {
	const line = lines[i].trim();
	if (line === '') continue;
	const values = line.split(',');
	const obj: Record<string, string> = {};
	obj['Name'] = values[0];
	obj['Age'] = values[1];
	obj['Email'] = values[2];
	obj['Score'] = values[3];
	rows.push(obj);
}

// Convert types
const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = rows.map(row => ({
	Name: row['Name'],
	Age: parseInt(row['Age'], 10),
	Email: row['Email'],
	Score: parseFloat(row['Score']),
}));

console.log(JSON.stringify(result, null, 2));