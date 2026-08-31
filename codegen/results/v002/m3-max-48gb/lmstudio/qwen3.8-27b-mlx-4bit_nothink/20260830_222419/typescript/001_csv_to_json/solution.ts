<tool_call>
<function=write_file>
<parameter=path>
main.ts
</parameter>
<parameter=content>
import * as fs from "fs";
import * as path from "path";

const inputPath = path.join("input", "data.csv");
const content = fs.readFileSync(inputPath, "utf-8");

const lines = content.split(/\r?\n/).filter((line) => line.trim() !== "");

if (lines.length === 0) {
  process.stdout.write("[]\n");
  process.exit(0);
}

const header = lines[0].split(",");

// Map header indices to canonical field names
const headerToField: Record<string, string> = {
  Name: "Name",
  Age: "Age",
  Email: "Email",
  Score: "Score",
};

const results: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(",");
  const obj: Record<string, string | number> = {};

  for (let j = 0; j < header.length; j++) {
    const colName = header[j].trim();
    const canonical = headerToField[colName];
    if (!canonical) continue;

    const raw = (values[j] ?? "").trim();

    if (canonical === "Age") {
      obj.Age = parseInt(raw, 10);
    } else if (canonical === "Score") {
      obj.Score = parseFloat(raw);
    } else {
      obj[canonical] = raw;
    }
  }

  results.push(obj as { Name: string; Age: number; Email: string; Score: number });
}

process.stdout.write(JSON.stringify(results, null, 2) + "\n");

</parameter>
</function>
</tool_call>