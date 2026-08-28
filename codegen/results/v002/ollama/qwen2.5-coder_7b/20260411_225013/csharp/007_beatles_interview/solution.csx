using System;
using System.Collections.Generic;
using System.Linq;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');
var data = lines.Skip(1).Select(line => line.Split(',')).ToList();

var result = new List<Dictionary<string, object>>();

foreach (var row in data)
{
    var person = new Dictionary<string, object>
    {
        { "FirstName", row[0].Split(' ')[0] },
        { "LastName", row[0].Split(' ')[1] },
        { "Birthday", DateTime.Parse(row[1]).ToString("yyyy-MM-dd") },
        { "Age", DateTime.Parse("2025-07-01").Year - DateTime.Parse(row[1]).Year }
    };

    var relatives = new List<Dictionary<string, object>>();

    if (!string.IsNullOrEmpty(row[4])) // Father
    {
        relatives.Add(new Dictionary<string, object>
        {
            { "FirstName", row[4].Split(' ')[0] },
            { "LastName", row[4].Split(' ')[1] },
            { "Relationship", "Father" }
        });
    }

    if (!string.IsNullOrEmpty(row[5])) // Mother
    {
        relatives.Add(new Dictionary<string, object>
        {
            { "FirstName", row[5].Split(' ')[0] },
            { "LastName", row[5].Split(' ')[1] },
            { "Relationship", "Mother" }
        });
    }

    if (!string.IsNullOrEmpty(row[6])) // Brother
    {
        relatives.Add(new Dictionary<string, object>
        {
            { "FirstName", row[6].Split(' ')[0] },
            { "LastName", row[6].Split(' ')[1] },
            { "Relationship", "Brother" }
        });
    }

    if (!string.IsNullOrEmpty(row[7])) // Sister
    {
        relatives.Add(new Dictionary<string, object>
        {
            { "FirstName", row[7].Split(' ')[0] },
            { "LastName", row[7].Split(' ')[1] },
            { "Relationship", "Sister" }
        });
    }

    person["Relatives"] = relatives;

    result.Add(person);
}

JsonSerializer.SerializeToUtf8String(result, Console.OpenStandardOutput(), new JsonSerializerOptions { WriteIndented = false });