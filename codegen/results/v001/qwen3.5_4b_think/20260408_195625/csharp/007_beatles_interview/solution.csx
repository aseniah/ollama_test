using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

// Read the expected format to understand the structure
string expectedPath = "input/expected_format.json";
string expected = JsonNode.Parse(File.ReadAllText(expectedPath));

// Get the expected structure to infer transformation rules
var expectedArray = expected.Value.AsJsonArray();
var expectedFields = expectedArray[0].Value?.PropertyNames();

// Create mapping from expected fields to input data
// Read input CSV
string inputPath = "input/input.csv";
string content = File.ReadAllText(inputPath);

// Parse CSV
var lines = content.Split('\n');
var headers = lines[0].Split(',').Select(s => s.Trim()).ToList();
var rowCount = lines.Length - 1;

// Data array to process
var data = new [rowCount]();

// Parse data with appropriate date format
var inputJson = new JsonArray();

// Process each row (excluding header)
for (int i = 1; i <= rowCount; i++)
{
    var line = lines[i];
    var fields = line.Split(',');
    var record = new object();
    
    // Map fields based on expected structure
    // Assuming input has: name, birthdate fields
    if (fields.Length >= 2)
    {
        record["name"] = fields[0].Trim();
        
        // Try to parse date - common formats
        string birthDate = fields[1].Trim();
        DateTime birthDateParsed;
        
        if (DateTime.TryParseExact(birthDate, "yyyy-MM-dd", null, System.Globalization.DateTimeStyles.None, out birthDateParsed))
        {
            record["birthDate"] = birthDateParsed;
        }
        else if (DateTime.TryParseExact(birthDate, "dd/MM/yyyy", null, System.Globalization.DateTimeStyles.None, out birthDateParsed))
        {
            record["birthDate"] = birthDateParsed;
        }
        else if (DateTime.TryParseExact(birthDate, "MM/dd/yyyy", null, System.Globalization.DateTimeStyles.None, out birthDateParsed))
        {
            record["birthDate"] = birthDateParsed;
        }
        else
        {
            // Try parsing without exact format
            DateTime.TryParse(birthDate, System.DateTimeStyles.None, null, out birthDateParsed);
            if (birthDateParsed.Year >= 1800 && birthDateParsed.Year <= 2100)
            {
                record["birthDate"] = birthDateParsed;
            }
        }
    }
    
    inputJson.Add(new JsonNode(record));
}

// Calculate ages as of July 1, 2025
var referenceDate = new DateTime(2025, 7, 1);
var outputJson = new JsonArray();

for (int i = 0; i < inputJson.Count; i++)
{
    JsonNode row = inputJson[i];
    string name = row.GetValue("name")?.Value<string>() ?? "Unknown";
    DateTime birthDate = row.GetValue("birthDate")?.Value<DateTime>() ?? DateTime.MinValue;
    
    // Calculate age
    int age = referenceDate.Year - birthDate.Year;
    
    // Adjust for birthday
    if (birthDate.Month > referenceDate.Month || 
        (birthDate.Month == referenceDate.Month && birthDate.Day > referenceDate.Day))
    {
        age--;
    }
    
    var record = new object[]
    {
        ("name", name),
        ("age", age)
    };
    
    outputJson.Add(new JsonNode(record));
}

// Output JSON array to stdout
Console.WriteLine(JsonSerializer.Serialize(outputJson, new JsonSerializerOptions
{
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
    DefaultValueHandling = DefaultValueHandling.Include
}));