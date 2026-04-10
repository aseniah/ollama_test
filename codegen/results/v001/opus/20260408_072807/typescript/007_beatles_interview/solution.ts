import * as fs from "fs";
import * as path from "path";

const csvPath = path.join(process.cwd(), "input", "input.csv");
const raw = fs.readFileSync(csvPath, "utf-8").trim();
const lines = raw.split("\n");
const headers = lines[0].split(",");

const REFERENCE_DATE = new Date(2025, 6, 1); // July 1, 2025

function calcAge(birthday: Date, died: Date | null): number {
  const end = died ?? REFERENCE_DATE;
  let age = end.getFullYear() - birthday.getFullYear();
  const m = end.getMonth() - birthday.getMonth();
  if (m < 0 || (m === 0 && end.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function parseName(full: string): { FirstName: string; LastName: string } {
  const parts = full.trim().split(" ");
  if (parts.length === 1) return { FirstName: parts[0], LastName: "" };
  return { FirstName: parts[0], LastName: parts[parts.length - 1] };
}

function parseDate(s: string): Date | null {
  if (!s || s === "null") return null;
  const [m, d, y] = s.split("/").map(Number);
  return new Date(y, m - 1, d);
}

const relationshipCols = ["Father", "Mother", "Brother", "Sister"];

const result = lines.slice(1).map((line) => {
  const values = line.split(",");
  const row: Record<string, string> = {};
  headers.forEach((h, i) => {
    row[h] = values[i] ?? "";
  });

  const { FirstName, LastName } = parseName(row["Name"]);
  const birthday = parseDate(row["Birthday"])!;
  const died = parseDate(row["Died"]);
  const age = calcAge(birthday, died);

  const birthdayStr = [
    birthday.getFullYear(),
    String(birthday.getMonth() + 1).padStart(2, "0"),
    String(birthday.getDate()).padStart(2, "0"),
  ].join("-");

  const relatives = relationshipCols
    .filter((rel) => row[rel] && row[rel] !== "null")
    .map((rel) => {
      const { FirstName: rFirst, LastName: rLast } = parseName(row[rel]);
      return { FirstName: rFirst, LastName: rLast, Relationship: rel };
    });

  return { FirstName, LastName, Birthday: birthdayStr, Age: age, Relatives: relatives };
});

console.log(JSON.stringify(result, null, 2));
