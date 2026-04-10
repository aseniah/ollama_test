import * as fs from "fs";

const filePath = process.argv[2] || "input/data.csv";

const content = fs.readFileSync(filePath, "utf8");
const lines = content.trim().split("\n");
if (lines.length < 1) {
    throw new Error("Empty file or missing header");
}

// Assuming the header is: Name,Age,Email,Score
const [header] = lines[0].split(",");

const rows = [];
for (let i = 1; i < lines.length; i++) {
    const fields = lines[i].split(",");
    if (fields.length !== 4) continue;

    const name = fields[0].trim();
    const age = parseInt(fields[1], 10);
    const email = fields[2].trim();
    const score = parseFloat(fields[3]);

    rows.push({ name, age, email, score });
}

console.log(JSON.stringify(rows, null, 2));