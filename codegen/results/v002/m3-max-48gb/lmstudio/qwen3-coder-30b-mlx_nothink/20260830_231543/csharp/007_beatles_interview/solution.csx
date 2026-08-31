using System;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');
var people = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    var values = lines[i].Split(',');
    var person = new JsonObject();
    
    var nameParts = values[0].Split(' ', StringSplitOptions.RemoveEmptyEntries);
    person["FirstName"] = nameParts[0];
    person["LastName"] = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    var birthday = DateTime.ParseExact(values[1], "M/d/yyyy", CultureInfo.InvariantCulture);
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    
    var age = CalculateAge(birthday, new DateTime(2025, 7, 1));
    person["Age"] = age;
    
    var relatives = new JsonArray();
    var relativeNames = new[] { values[3], values[4], values[5], values[6] };
    var relationships = new[] { "Father", "Mother", "Brother", "Sister" };
    
    for (int j = 0; j < relativeNames.Length; j++)
    {
        if (!string.IsNullOrEmpty(relativeNames[j]))
        {
            var relative = new JsonObject();
            var relativeParts = relativeNames[j].Split(' ', StringSplitOptions.RemoveEmptyEntries);
            relative["FirstName"] = relativeParts[0];
            relative["LastName"] = relativeParts.Length > 1 ? string.Join(" ", relativeParts.Skip(1)) : "";
            relative["Relationship"] = relationships[j];
            relatives.Add(relative);
        }
    }
    
    person["Relatives"] = relatives;
    people.Add(person);
}

Console.WriteLine(people.ToJsonString());

int CalculateAge(DateTime birthDate, DateTime referenceDate)
{
    int age = referenceDate.Year - birthDate.Year;
    if (birthDate.AddYears(age) > referenceDate)
        age--;
    return age;
}