using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the CSV file
string[] lines = File.ReadAllLines("input/input.csv");

// Parse the header
string[] headers = lines[0].Split(',');

// Create a list to hold the person objects
var people = new JsonArray();

// Start from line 1 to skip the header
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    string[] values = line.Split(',');
    
    if (values.Length != headers.Length)
        continue;

    // Create a dictionary for easy access by header name
    var row = new Dictionary<string, string>();
    for (int j = 0; j < headers.Length; j++)
    {
        row[headers[j].Trim()] = values[j].Trim();
    }

    string fullName = row["Name"];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    string birthdayStr = row["Birthday"];
    
    // Parse birthday
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);

    // Calculate age as of July 1, 2025
    DateTime referenceDate = new DateTime(2025, 7, 1);
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    // Format birthday as YYYY-MM-DD
    string formattedBirthday = birthday.ToString("yyyy-MM-dd");

    // Build the person object
    var person = new JsonObject();
    person["FirstName"] = firstName;
    person["LastName"] = lastName;
    person["Birthday"] = formattedBirthday;
    person["Age"] = age;

    // Build relatives array
    var relatives = new JsonArray();
    
    string father = row["Father"];
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        string[] fatherParts = father.Split(' ');
        var relative = new JsonObject();
        relative["FirstName"] = fatherParts[0];
        relative["LastName"] = fatherParts[fatherParts.Length - 1];
        relative["Relationship"] = "Father";
        relatives.Add(relative);
    }

    string mother = row["Mother"];
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        string[] motherParts = mother.Split(' ');
        var relative = new JsonObject();
        relative["FirstName"] = motherParts[0];
        relative["LastName"] = motherParts[motherParts.Length - 1];
        relative["Relationship"] = "Mother";
        relatives.Add(relative);
    }

    string brother = row["Brother"];
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        string[] brotherParts = brother.Split(' ');
        var relative = new JsonObject();
        relative["FirstName"] = brotherParts[0];
        relative["LastName"] = brotherParts[brotherParts.Length - 1];
        relative["Relationship"] = "Brother";
        relatives.Add(relative);
    }

    string sister = row["Sister"];
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        string[] sisterParts = sister.Split(' ');
        var relative = new JsonObject();
        relative["FirstName"] = sisterParts[0];
        relative["LastName"] = sisterParts[sisterParts.Length - 1];
        relative["Relationship"] = "Sister";
        relatives.Add(relative);
    }

    person["Relatives"] = relatives;
    people.Add(person);
}

// Set output formatting
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    DefaultIgnoreCondition = System.Text.Json.Serialization.JsonIgnoreCondition.WhenWritingNull
};

// Write the JSON array to stdout
Console.WriteLine(people.ToJsonString(options));