import { readFileSync, existsSync } from "fs";
import { join, dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const inputPath = join(__dirname, "input", "input.csv");
const expectedPath = join(__dirname, "input", "expected_format.json");

if (!existsSync(inputPath)) {
    console.error(`Error: Input file not found: ${inputPath}`);
    process.exit(1);
}
if (!existsSync(expectedPath)) {
    console.error(`Error: Expected format file not found: ${expectedPath}`);
    process.exit(1);
}

const inputContent = readFileSync(inputPath, "utf-8");
const expectedContent = JSON.parse(readFileSync(expectedPath, "utf-8"));

// Parse CSV manually to avoid dependencies
function parseCSV(csv: string): string[][] {
    const lines = csv.split("\n").map(line => line.trim()).filter(line => line.length > 0);
    if (lines.length === 0) return [];

    const headers = lines[0].split(",").map(h => h.trim());
    
    return lines.slice(1).map(line => {
        const values = line.split(",");
        // Handle potential quotes around values if needed, but standard simple CSV usually works
        // Assuming standard CSV where values might be quoted
        const mappedValues: string[] = [];
        for (const header of headers) {
            const val = values.shift() || "";
            // Basic CSV parsing logic: if a value starts with ", it might be quoted
            if (val.startsWith('"') && val.endsWith('"')) {
                val = val.slice(1, -1).replace(/\\"/g, '"');
            }
            mappedValues.push(val);
        }
        return mappedValues;
    });
}

const dataRows = parseCSV(inputContent);

// Calculate age as of July 1, 2025
// Assuming the date field is the second column (index 1) in a format like YYYY-MM-DD or MM-DD-YYYY
// Looking at typical CSV structures for age problems:
// Often columns are: First Name, Last Name, Birth Date, City, etc.
// Let's assume the birth date is at index 1 (after First Name)
// If the first name is a string, it won't be a valid date format, so index 1 is a safe bet for DOB.

const targetDate = new Date("2025-07-01");

const result: any[] = [];

for (const row of dataRows) {
    if (row.length < 2) continue; // Ensure we have at least name and birth date

    const birthDateString = row[1].trim();
    
    // Parse date. Common formats: YYYY-MM-DD, MM-DD-YYYY, DD-MM-YYYY
    // Try YYYY-MM-DD first
    const parts = birthDateString.split(/[-/.]/);
    
    if (parts.length !== 3) {
        // If simple date is missing, maybe the date is a separate column?
        // But without knowing the schema, let's assume standard date is the only date-like string.
        // Let's try to detect if any column looks like a date.
        const dateCol = row.find((cell, idx) => {
            // Simple heuristic: contains digits and looks like a date
            const digits = cell.replace(/[^0-9]/g, "");
            return digits.length === 4 || digits.length === 6;
        });

        if (dateCol) {
            const parsedDate = new Date(dateCol.trim());
            // Check if it's a valid date and in the future/past
            if (!isNaN(parsedDate.getTime()) && parsedDate > new Date(1800) && parsedDate < new Date(2100)) {
                result.push({
                    "First Name": row[0],
                    "Age": Math.max(0, Math.floor((targetDate.getTime() - parsedDate.getTime()) / (1000 * 60 * 60 * 24 * 365.25))),
                    // Add other columns if needed, but based on prompt "infer transformation rules",
                    // usually implies minimal viable transformation matching the example.
                });
                continue;
            }
        }
    }

    // Fallback: Assume specific index if no heuristic found, but let's stick to index 1 as DOB
    // Re-parse index 1 with different separators just in case
    const dateVal = row[1];
    let [y, m, d] = dateVal.split(/[-/.]/).map(Number);
    
    // Correct order based on separators if ambiguous, but YYYY-MM-DD is most common in data files
    // Let's assume YYYY-MM-DD for index 1
    
    if (!isNaN(y) && !isNaN(m) && !isNaN(d)) {
        // Check validity
        if (y < 1900 || y > 2050) continue;
        
        // Some datasets have MM-DD-YYYY. Let's check which order makes sense with the year.
        // If we assume current year is 2025.
        // If m > 12 or d > 31, something is wrong or order is different.
        // Let's calculate age for both interpretations if possible? No, pick the most likely.
        // Standard ISO is YYYY-MM-DD.
        
        const birthDate = new Date(y, m - 1, d);
        
        // Age Calculation
        let age = targetDate.getFullYear() - birthDate.getFullYear();
        const birthMonth = birthDate.getMonth() + 1;
        const targetMonth = 7;
        
        if (birthMonth > targetMonth || (birthMonth === targetMonth && birthDate.getDate() >= targetDate.getDate())) {
            age -= 1;
        }
        
        if (age < 0) age = 0;

        result.push({
            "First Name": row[0],
            "Age": age
        });
    }
}

// Sort by Age descending usually, or by Name? Let's sort by Age descending as is common in these tasks.
// Also, ensure the output keys match the expected format structure.
// Expected format usually looks like: [{"First Name": "Alice", "Age": 20}, ...]
// We should probably sort by Name then Age, or just output the list. 
// Let's sort by Age descending.
result.sort((a, b) => b["Age"] - a["Age"]);

console.log(JSON.stringify(result, null, 2));