const fs = require('fs');
const path = require('path');

const inputPath = path.join(__dirname, 'input', 'data.csv');
let content;
try {
	content = fs.readFileSync(inputPath, 'utf8');
} catch (e) {
	process.exitCode = 1;
	process.exit();
}

const lines = content.trim().split(/\r?\n/);
if (lines.length === 0) {
	process.exitCode = 0;
}

const rows = [];

for (const line of lines.slice(1)) {
	const parts = line.split(',');
	
	const name = parts[0];
	
	const agePart = parts[1];
	const age = parseInt(agePart, 10);
	
	const email = parts[2];
	
	const scorePart = parts[3];
	const score = parseFloat(scorePart);
	
	rows.push({
		Name: name,
		Age: age,
		Email: email,
		Score: score,
	});
}

console.log(JSON.stringify(rows));