#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Read CSV file
string csvPath = "input/input.csv";
string csvContent = File.ReadAllText(csvPath);
string[] lines = csvContent.Split(new[] { '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Parse CSV with proper handling of quoted fields
IEnumerable<string[]> ParseCSV(string[] lines)
{
    foreach (string line in lines)
    {
        List<string> fields = new List<string>();
        int pos = 0;
        bool inQuotes = false;
        StringBuilder field = new StringBuilder();
        
        while (pos < line.Length)
        {
            char c = line[pos];
            if (inQuotes)
            {
                if (c == '"')
                {
                    if (pos + 1 < line.Length && line[pos + 1] == '"')
                    {
                        field.Append('"');
                        pos += 2;
                    }
                    else
                    {
                        inQuotes = false;
                        pos++;
                    }
                }
                else
                {
                    field.Append(c);
                    pos++;
                }
            }
            else
            {
                if (c == '"')
                {
                    inQuotes = true;
                    pos++;
                }
                else if (c == ',')
                {
                    fields.Add(field.ToString());
                    field.Clear();
                    pos++;
                }
                else
                {
                    field.Append(c);
                    pos++;
                }
            }
        }
        fields.Add(field.ToString());
        yield return fields.ToArray();
    }
}

var csvRows = ParseCSV(lines).ToList();
if (csvRows.Count == 0)
{
    Console.WriteLine("[]");
    return;
}

// Parse header
string[] headers = csvRows[0];
for (int i = 0; i < headers.Length; i++)
{
    headers[i] = headers[i].Trim();
}

// Build JSON array
JsonArray jsonArray = new JsonArray();

for (int rowIdx = 1; rowIdx < csvRows.Count; rowIdx++)
{
    string[] values = csvRows[rowIdx];
    JsonArray jsonRow = new JsonArray();
    
    Dictionary<string, string> rowDict = new Dictionary<string, string>(StringComparer.OrdinalIgnoreCase);
    for (int i = 0; i < headers.Length && i < values.Length; i++)
    {
        string header = headers[i];
        string value = i < values.Length ? values[i].Trim() : "";
        rowDict[header] = value;
        
        // Handle JSON value conversion
        if (int.TryParse(value, out int intVal))
        {
            jsonRow.Add(JsonNode.Parse(intVal.ToString()));
        }
        else if (double.TryParse(value, out double doubleVal))
        {
            jsonRow.Add(JsonNode.Parse(doubleVal.ToString()));
        }
        else
        {
            jsonRow.Add(value);
        }
    }
    
    // Find date of birth field and calculate age
    string? dobHeader = null;
    foreach (string h in headers)
    {
        string lower = h.ToLowerInvariant();
        if (lower.Contains("dob") || lower.Contains("birthdate") || 
            lower.Contains("birth_date") || lower.Contains("birth") || 
            lower.Contains("dateofbirth") || lower.Contains("dateofbirth"))
        {
            dobHeader = h;
            break;
        }
    }
    
    if (dobHeader != null && rowDict.ContainsKey(dobHeader))
    {
        string dobValue = rowDict[dobHeader];
        if (DateTime.TryParse(dobValue, CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime birthDate))
        {
            int age = referenceDate.Year - birthDate.Year;
            if (referenceDate < birthDate.AddYears(age))
            {
                age--;
            }
            jsonRow.Add(dobHeader + "_age", age);
        }
    }
    
    jsonArray.Add(jsonRow);
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray));