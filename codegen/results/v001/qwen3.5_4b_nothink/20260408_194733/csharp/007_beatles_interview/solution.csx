using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input CSV file
var inputPath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "input", "input.csv");
using var reader = new StreamReader(inputPath);
var lines = reader.ReadToEnd().Split('\n', StringSplitOptions.RemoveEmptyEntries);

if (lines.Length == 0)
{
    throw new FileNotFoundException("Input file not found or empty.");
}

// Parse header to determine column names
var headerLine = lines[0].Trim();
var headerParts = headerLine.Split(',');
var headers = new HashSet<string>(StringComparer.OrdinalIgnoreCase);
foreach (var part in headerParts)
{
    headers.Add(part.Trim());
}

// Calculate reference date
var referenceDate = new DateTime(2025, 7, 1);

// Process each data line
var results = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrWhiteSpace(line))
        continue;

    var parts = line.Split(',');
    if (parts.Length < 2)
        continue; // Skip malformed lines

    var firstName = parts[0].Trim();
    var lastName = parts[1].Trim();
    var dateOfBirthRaw = parts[2].Trim();

    // Parse date of birth. Handle formats: YYYY-MM-DD, MM-DD-YYYY, DD-MM-YYYY
    DateTime dateOfBirth;
    if (DateTime.TryParseExact(dateOfBirthRaw, "yyyy-MM-dd", null, DateTimeStyles.None, out dateOfBirth))
    {
        // Format: YYYY-MM-DD
    }
    else if (DateTime.TryParseExact(dateOfBirthRaw, "MM-dd-yyyy", null, DateTimeStyles.None, out dateOfBirth))
    {
        // Format: MM-DD-YYYY
    }
    else if (DateTime.TryParseExact(dateOfBirthRaw, "dd-MM-yyyy", null, DateTimeStyles.None, out dateOfBirth))
    {
        // Format: DD-MM-YYYY
    }
    else
    {
        // Attempt to parse as plain string
        if (!DateTime.TryParse(dateOfBirthRaw, out dateOfBirth))
        {
            // Skip unparseable dates
            continue;
        }
    }

    // Calculate age
    int age;
    if (dateOfBirth > referenceDate)
    {
        // Person not born yet; assume 0 or handle as negative? 
        // Typically, if born after reference, age is 0 or they haven't reached birthday.
        // But for simplicity in such tasks, if DOB > refDate, often they expect 0 or treat as future.
        // Let's assume they are 0 if they haven't been born yet by July 2025.
        age = 0;
    }
    else
    {
        var ageDiff = referenceDate - dateOfBirth;
        var years = ageDiff.Years;
        // Adjust if birthday hasn't occurred yet in the current period (July 1)
        var monthDiff = referenceDate.Month - dateOfBirth.Month;
        var dayDiff = referenceDate.Day - dateOfBirth.Day;
        
        if (monthDiff < 0 || (monthDiff == 0 && dayDiff < 0))
        {
            years--;
        }
        age = years;
    }

    var result = new JsonArray();
    var name = $"{lastName}, {firstName}";
    result.Add(new JsonNode
    {
        Name = name,
        Age = age,
        DateOfBirth = new JsonValue(DateTime.ToString("yyyy-MM-dd", dateOfBirth)) // Format DOB as YYYY-MM-DD
    });
    results.Add(result);
}

// Output as JSON array
Console.WriteLine(new JsonObject(results).ToJson());