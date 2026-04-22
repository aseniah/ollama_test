using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the CSV file
string[] lines = File.ReadAllLines("input/input.csv");

// Skip the header line
List<string> dataLines = new List<string>();
for (int i = 1; i < lines.Length; i++)
{
    if (!string.IsNullOrWhiteSpace(lines[i]))
    {
        dataLines.Add(lines[i]);
    }
}

// Define the reference date for age calculation: July 1, 2025
DateTime referenceDate = new DateTime(2025, 7, 1);

// Parse the JSON structure from the expected format to understand the output
// But since we need to produce the JSON, we'll build it programmatically.

JsonArray result = new JsonArray();

foreach (string line in dataLines)
{
    string[] fields = line.Split(',');
    
    // Fields: Name, Birthday, Died, Father, Mother, Brother, Sister
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string fatherName = fields[3].Trim();
    string motherName = fields[4].Trim();
    string brotherName = fields[5].Trim();
    string sisterName = fields[6].Trim();
    
    // Parse the name into FirstName and LastName
    // The name seems to be "FirstName MiddleName LastName" or "FirstName LastName"
    // Let's split by space. The last part is LastName, the first part is FirstName.
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    // Parse the birthday
    // Format is M/d/yyyy or MM/dd/yyyy
    DateTime birthday;
    if (DateTime.TryParse(birthdayStr, out birthday))
    {
        // Calculate age as of July 1, 2025
        int age = referenceDate.Year - birthday.Year;
        if (referenceDate.Month < birthday.Month || 
            (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
        {
            age--;
        }
        
        // Create the main person object
        JsonObject person = new JsonObject
        {
            ["FirstName"] = firstName,
            ["LastName"] = lastName,
            ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
            ["Age"] = age
        };
        
        // Create the relatives array
        JsonArray relatives = new JsonArray();
        
        // Helper function to add a relative if the name is not null
        void AddRelative(string nameStr, string relationship)
        {
            if (nameStr != "null" && !string.IsNullOrWhiteSpace(nameStr))
            {
                string[] parts = nameStr.Split(' ');
                string relFirstName = parts[0];
                string relLastName = parts[parts.Length - 1];
                
                JsonObject relative = new JsonObject
                {
                    ["FirstName"] = relFirstName,
                    ["LastName"] = relLastName,
                    ["Relationship"] = relationship
                };
                relatives.Add(relative);
            }
        }
        
        AddRelative(fatherName, "Father");
        AddRelative(motherName, "Mother");
        AddRelative(brotherName, "Brother");
        AddRelative(sisterName, "Sister");
        
        person["Relatives"] = relatives;
        result.Add(person);
    }
}

// Output the JSON array
Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));