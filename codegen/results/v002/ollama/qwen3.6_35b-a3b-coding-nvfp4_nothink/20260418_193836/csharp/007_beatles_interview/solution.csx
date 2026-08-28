using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the CSV file
string[] lines = File.ReadAllLines("input/input.csv");

// Define the reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Parse the CSV
// Skip header line
List<JsonObject> people = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    string[] parts = line.Split(',');
    // Expected format: Name, Birthday, Died, Father, Mother, Brother, Sister
    
    string fullName = parts[0].Trim();
    string birthdayStr = parts[1].Trim();
    string diedStr = parts[2].Trim(); // Not used for output but good to know
    string fatherName = parts[3].Trim();
    string motherName = parts[4].Trim();
    string brotherName = parts[5].Trim();
    string sisterName = parts[6].Trim();

    // Parse Name: "FirstName MiddleName LastName" or "FirstName LastName"
    // The expected output only has FirstName and LastName.
    // Looking at the data:
    // "John Winston Lennon" -> FirstName: John, LastName: Lennon
    // "James Paul McCartney" -> FirstName: James, LastName: McCartney
    // "Ringo Starr" -> FirstName: Ringo, LastName: Starr
    // "George Harrison" -> FirstName: George, LastName: Harrison
    
    string[] nameParts = fullName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday: "M/D/YYYY" or "MM/D/YYYY" etc.
    DateTime birthDate;
    if (!DateTime.TryParse(birthdayStr, out birthDate))
    {
        // Try parsing with a specific format if needed, but TryParse usually handles M/D/YYYY
        // Let's be safe and use a parser if it fails
        birthDate = DateTime.MinValue;
    }

    // Calculate Age as of July 1, 2025
    int age = 0;
    if (birthDate != DateTime.MinValue)
    {
        age = referenceDate.Year - birthDate.Year;
        if (referenceDate < birthDate.AddYears(age))
        {
            age--;
        }
    }

    // Create the person object
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthDate.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    // Add Father
    if (fatherName != null && fatherName != "null" && !string.IsNullOrWhiteSpace(fatherName))
    {
        string[] fParts = fatherName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        string fFirst = fParts[0];
        string fLast = fParts[fParts.Length - 1];
        JsonObject father = new JsonObject
        {
            ["FirstName"] = fFirst,
            ["LastName"] = fLast,
            ["Relationship"] = "Father"
        };
        person["Relatives"].AsArray().Add(father);
    }

    // Add Mother
    if (motherName != null && motherName != "null" && !string.IsNullOrWhiteSpace(motherName))
    {
        string[] mParts = motherName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        string mFirst = mParts[0];
        string mLast = mParts[mParts.Length - 1];
        JsonObject mother = new JsonObject
        {
            ["FirstName"] = mFirst,
            ["LastName"] = mLast,
            ["Relationship"] = "Mother"
        };
        person["Relatives"].AsArray().Add(mother);
    }

    // Add Brother
    if (brotherName != null && brotherName != "null" && !string.IsNullOrWhiteSpace(brotherName))
    {
        string[] bParts = brotherName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        string bFirst = bParts[0];
        string bLast = bParts[bParts.Length - 1];
        JsonObject brother = new JsonObject
        {
            ["FirstName"] = bFirst,
            ["LastName"] = bLast,
            ["Relationship"] = "Brother"
        };
        person["Relatives"].AsArray().Add(brother);
    }

    // Add Sister
    if (sisterName != null && sisterName != "null" && !string.IsNullOrWhiteSpace(sisterName))
    {
        string[] sParts = sisterName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        string sFirst = sParts[0];
        string sLast = sParts[sParts.Length - 1];
        JsonObject sister = new JsonObject
        {
            ["FirstName"] = sFirst,
            ["LastName"] = sLast,
            ["Relationship"] = "Sister"
        };
        person["Relatives"].AsArray().Add(sister);
    }

    people.Add(person);
}

// Serialize to JSON
JsonSerializerOptions options = new JsonSerializerOptions
{
    WriteIndented = true
};

string json = JsonSerializer.Serialize(people, options);
Console.WriteLine(json);