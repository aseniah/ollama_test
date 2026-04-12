import fs from "fs";

// Read input file
const rawData = fs.readFileSync("input/data.json", "utf-8");
const data: unknown[] = JSON.parse(rawData);

// Filter records where active is true and age >= 30, then sort by name ascending
const filtered = data.filter<unknown>(record => 
    Boolean(record.active) && Number(record.age) >= 30
);

filtered.sort((a, b) => {
    const nameA = String(a.name);
    const nameB = String(b.name);
    return nameA.localeCompare(nameB);
});

// Output result as JSON to stdout
console.log(JSON.stringify(filtered, null, 2));