using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var result = new JsonArray();
var refDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    string[] nameParts = parts[0].Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    string[] bdayParts = parts[1].Split('/');
    int year = int.Parse(bdayParts[2]);
    int month = int.Parse(bdayParts[0]);
    int day = int.Parse(bdayParts[1]);
    string formattedBirthday = $"{year:D4}-{month:D2}-{day:D2}";

    int age = refDate.Year - year;
    if (refDate.Month < month || (refDate.Month == month && refDate.Day < day))
        age--;

    var person = new JsonObject();
    person["FirstName"] = firstName;
    person["LastName"] = lastName;
    person["Birthday"] = formattedBirthday;
    person["Age"] = age;
    person["Relatives"] = new JsonArray();

    string[] relNames = { parts[3].Trim(), parts[4].Trim(), parts[5].Trim(), parts[6].Trim() };
    string[] relTypes = { "Father", "Mother", "Brother", "Sister" };

    for (int j = 0; j < relNames.Length; j++)
    {
        if (relNames[j] != "null")
        {
            string[] relParts = relNames[j].Split(' ');
            var relObj = new JsonObject();
            relObj["FirstName"] = relParts[0];
            relObj["LastName"] = relParts[relParts.Length - 1];
            relObj["Relationship"] = relTypes[j];
            ((JsonArray)person["Relatives"])!.Add(relObj);
        }
    }

    result.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));