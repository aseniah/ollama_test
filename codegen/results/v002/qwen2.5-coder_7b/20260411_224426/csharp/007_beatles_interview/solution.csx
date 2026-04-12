using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;

var lines = File.ReadAllLines("input/input.csv");

var people = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    
    var person = new Dictionary<string, object>
    {
        { "FirstName", parts[0].Split(' ').First() },
        { "LastName", parts[0].Split(' ').Last() },
        { "Birthday", DateTime.Parse(parts[1], CultureInfo.InvariantCulture) },
        { "Age", DateTime.Parse("7/1/2025", CultureInfo.InvariantCulture).Year - (int)DateTime.Parse(parts[1], CultureInfo.InvariantCulture).ToString("yyyy") }
    };

    var relatives = new List<Dictionary<string, string>>
    {
        AddRelative(parts[3], parts[4], "Father"),
        AddRelative(parts[3], parts[4], "Mother"),
        AddRelative(parts[5], null, "Brother"),
        AddRelative(parts[6], null, "Sister")
    };

    person["Relatives"] = relatives.Where(r => r != null).ToList();

    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people));

Dictionary<string, string> AddRelative(string firstName, string lastName, string relationship)
{
    if (string.IsNullOrEmpty(firstName) || string.IsNullOrEmpty(lastName))
        return null;

    return new Dictionary<string, string>
    {
        { "FirstName", firstName },
        { "LastName", lastName },
        { "Relationship", relationship }
    };
}