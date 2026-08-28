using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');

var people = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var person = new Dictionary<string, object>();
    var row = lines[i].Split(',');

    for (int j = 0; j < header.Length; j++)
    {
        string key = header[j];
        string value = row[j];

        if (key == "Birthday" || key == "Died")
        {
            DateTime date;
            if (DateTime.TryParseExact(value, "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out date))
            {
                value = date.ToString("yyyy-MM-dd");
            }
        }

        person[key] = value;
    }

    var relatives = new List<Dictionary<string, string>>();
    if (!string.IsNullOrEmpty(person["Father"] as string))
    {
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", person["Father"] as string },
            { "LastName", person["Mother"] is null ? "" : ((string)person["Mother"]).Split(' ')[0] },
            { "Relationship", "Father" }
        });
    }
    if (!string.IsNullOrEmpty(person["Mother"] as string))
    {
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", person["Mother"] is null ? "" : ((string)person["Mother"]).Split(' ')[0] },
            { "LastName", person["Father"] is null ? "" : ((string)person["Father"]).Split(' ')[1] },
            { "Relationship", "Mother" }
        });
    }
    if (!string.IsNullOrEmpty(person["Brother"] as string))
    {
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", person["Brother"] as string },
            { "LastName", person["Father"] is null ? "" : ((string)person["Father"]).Split(' ')[1] },
            { "Relationship", "Brother" }
        });
    }
    if (!string.IsNullOrEmpty(person["Sister"] as string))
    {
        relatives.Add(new Dictionary<string, string>
        {
            { "FirstName", person["Sister"] as string },
            { "LastName", person["Father"] is null ? "" : ((string)person["Father"]).Split(' ')[1] },
            { "Relationship", "Sister" }
        });
    }

    if (relatives.Count > 0)
    {
        person["Relatives"] = relatives;
    }

    DateTime birthday = DateTime.ParseExact(person["Birthday"] as string, "yyyy-MM-dd", CultureInfo.InvariantCulture);
    int age = DateTime.Compare(new DateTime(2025, 7, 1), birthday) < 0 ? (int)(DateTime.DaysInMonth(2025, 7) + birthday.DayOfYear - DateTime.ParseExact(person["Birthday"] as string, "yyyy-MM-dd", CultureInfo.InvariantCulture).DayOfYear) : (int)((new DateTime(2025, 7, 1) - birthday).Days / 365);
    person["Age"] = age;

    people.Add(person);
}

var jsonOptions = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(people, jsonOptions));