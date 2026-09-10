using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");

// Read expected format to confirm structure (not strictly needed but we follow it)
var result = new JsonArray();

DateTime referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++) // skip header
{
    var parts = lines[i].Split(',');
    string fullName = parts[0].Trim();
    string birthdayStr = parts[1].Trim();
    string diedStr = parts[2].Trim();
    string fatherStr = parts[3].Trim();
    string motherStr = parts[4].Trim();
    string brotherStr = parts[5].Trim();
    string sisterStr = parts[6].Trim();

    // Split name: first word = FirstName, last word = LastName
    var nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse birthday MM/DD/YYYY
    var bdayParts = birthdayStr.Split('/');
    int bdayMonth = int.Parse(bdayParts[0]);
    int bdayDay = int.Parse(bdayParts[1]);
    int bdayYear = int.Parse(bdayParts[2]);
    string birthdayFormatted = $"{bdayYear:D4}-{bdayMonth:D2}-{bdayDay:D2}";

    // Calculate age
    int age;
    DateTime targetDate;
    bool isDead = diedStr != "null";
    
    if (isDead)
    {
        var diedParts = diedStr.Split('/');
        int diedMonth = int.Parse(diedParts[0]);
        int diedDay = int.Parse(diedParts[1]);
        int diedYear = int.Parse(diedParts[2]);
        targetDate = new DateTime(diedYear, diedMonth, diedDay);
    }
    else
    {
        targetDate = referenceDate;
    }

    int years = targetDate.Year - bdayYear;
    // Check if birthday has passed by target date
    if (targetDate.Month < bdayMonth || (targetDate.Month == bdayMonth && targetDate.Day < bdayDay))
    {
        years--;
    }
    age = years;

    // Build relatives
    var relatives = new JsonArray();

    void AddRelative(string relName, string relationship)
    {
        if (relName == "null") return;
        var rParts = relName.Split(' ');
        var rel = new JsonObject
        {
            ["FirstName"] = rParts[0],
            ["LastName"] = rParts[rParts.Length - 1],
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }

    AddRelative(fatherStr, "Father");
    AddRelative(motherStr, "Mother");
    AddRelative(brotherStr, "Brother");
    AddRelative(sisterStr, "Sister");

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    result.Add(person);
}

Console.Write(result.ToJsonString(new JsonWriterOptions { Indented = true }));