using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');

var people = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    // Parse name
    var nameParts = fields[0].Trim().Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];
    
    // Parse birthday: M/D/YYYY -> YYYY-MM-DD
    var bdParts = fields[1].Trim().Split('/');
    var birthYear = int.Parse(bdParts[2]);
    var birthMonth = int.Parse(bdParts[0]);
    var birthDay = int.Parse(bdParts[1]);
    var birthdayStr = $"{birthYear:D4}-{birthMonth:D2}-{birthDay:D2}";
    
    // Calculate age
    int age;
    var diedStr = fields[2].Trim();
    int refYear, refMonth, refDay;
    
    if (diedStr == "null")
    {
        refYear = 2025; refMonth = 7; refDay = 1;
    }
    else
    {
        var diedParts = diedStr.Split('/');
        refMonth = int.Parse(diedParts[0]);
        refDay = int.Parse(diedParts[1]);
        refYear = int.Parse(diedParts[2]);
    }
    
    age = refYear - birthYear;
    if (refMonth < birthMonth || (refMonth == birthMonth && refDay < birthDay))
    {
        age--;
    }
    
    // Build relatives
    var relatives = new JsonArray();
    var relCols = new[] { 3, 4, 5, 6 };
    var relNames = new[] { "Father", "Mother", "Brother", "Sister" };
    
    for (int j = 0; j < 4; j++)
    {
        var relStr = fields[relCols[j]].Trim();
        if (relStr == "null") continue;
        
        var relParts = relStr.Split(' ');
        var relObj = new JsonObject
        {
            ["FirstName"] = relParts[0],
            ["LastName"] = relParts[relParts.Length - 1],
            ["Relationship"] = relNames[j]
        };
        relatives.Add(relObj);
    }
    
    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    people.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(people.ToJsonString(options));