import { readFileSync } from "node:fs";

const rawData = JSON.parse(readFileSync("input/data.json", "utf-8"));

const filtered = rawData.filter((record: any) => record.active && record.age >= 30);

const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(sorted, null, 2));