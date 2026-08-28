import * as fs from "fs";
import * as path from "path";

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

function parseDate(dateStr: string): { year: number; month: number; day: number } {
  const parts = dateStr.split("/").map(Number);
  return { month: parts[0], day: parts[1], year: parts[2] };
}

function formatDate(dateStr: string): string {
  const { year, month, day } = parseDate(dateStr);
  return `${year}-${String(month).padStart(2, "0")}-${String(day).padStart(2, "0")}`;
}

function calculateAge(birthday: { year: number; month: number; day: number }, reference: { year: number; month: number; day: number }): number {
  let age = reference.year - birthday.year;
  if (reference.month < birthday.month || (reference.month === birthday.month && reference.day < birthday.day)) {
    age--;
  }
  return age;
}

function main(): void {
  const csvPath = path.resolve("input/input.csv");
  const content = fs.readFileSync(csvPath, "utf-8").trim();
  const lines = content.split("\n");
  const headers = lines[0].split(",");
  const rows = lines.slice(1).map((line) => {
    const cols = line.split(",");
    const obj: Record<string, string> = {};
    headers.forEach((h, i) => {
      obj[h.trim()] = (cols[i] || "").trim();
    });
    return obj;
  });

  const referenceDate = { year: 2025, month: 7, day: 1 };

  const result: Person[] = rows.map((row) => {
    const fullName = row["Name"];
    const nameParts = fullName.split(" ");
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];

    const birthday = parseDate(row["Birthday"]);
    const died = row["Died"] !== "null" ? row["Died"] : null;

    let age: number;
    if (died) {
      const deathDate = parseDate(died);
      age = calculateAge(birthday, deathDate);
    } else {
      age = calculateAge(birthday, referenceDate);
    }

    const relatives: Relative[] = [];
    const relMap: Array<{ key: string; relationship: string }> = [
      { key: "Father", relationship: "Father" },
      { key: "Mother", relationship: "Mother" },
      { key: "Brother", relationship: "Brother" },
      { key: "Sister", relationship: "Sister" },
    ];

    for (const { key, relationship } of relMap) {
      const val = row[key];
      if (val && val !== "null") {
        const parts = val.split(" ");
        relatives.push({
          FirstName: parts[0],
          LastName: parts[parts.length - 1],
          Relationship: relationship,
        });
      }
    }

    return {
      FirstName: firstName,
      LastName: lastName,
      Birthday: formatDate(row["Birthday"]),
      Age: age,
      Relatives: relatives,
    };
  });

  process.stdout.write(JSON.stringify(result, null, 2) + "\n");
}

main();