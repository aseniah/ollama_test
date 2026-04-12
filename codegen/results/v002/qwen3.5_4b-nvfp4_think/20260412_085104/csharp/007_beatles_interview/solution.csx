using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json.Nodes;

var currentYear = 2025;
var currentDate = new DateTime(currentYear, 7, 1);

var lines = File.ReadAllLines("input/input.csv");

var headers = lines.First().Split(',');
var persons = new List<Dictionary<string, string>>();

for (var i = 1; i < lines.Length; i++)
{
    var row = lines[i].Trim();
    if (row == "Name,Birthday,Died,Father,Mother,Brother,Sister") continue;

    var values = row.Split(',');
    var birthDate = new DateTime(int.Parse(values[1].Split('/').First().Trim()), 
                                 int.Parse(values[1].Split('/')[1].Trim()), 
                                 int.Parse(values[1].Split('/')[2].Trim()));
    
    var people = new Dictionary<string, string> {{ "FirstName", values[0] }, { "LastName", "" } };
    
    foreach (var namePair in headers.Skip(1).SelectMany(h => h.Split(',')))
    {
        if (namePair != null && !string.IsNullOrEmpty(namePair))
        {
            var field = namePair;
        }
    }
}

Console.Write("[\n");
int count = 0;
var totalRows = lines.Length - 1;

for (var i = 1; i < lines.Length; i++)
{
    var row = lines[i].Trim();
    if (row == "Name,Birthday,Died,Father,Mother,Brother,Sister") continue;
    
    var values = row.Split(',');
    var firstName = values[0];
    var birthParts = values[1].Split('/');
    var birthYear = int.Parse(birthParts[2]);
    var person = new Dictionary<string, string> { 
        {"FirstName", values[0]}, 
        {"LastName", ""}, 
        {"Birthday", $"{birthParts[0]}/{birthParts[1]}/{birthParts[2]}"} 
    };
    
    var age = currentYear - birthYear;
    if (age < 0) age = 0;
    
    person["Age"] = age.ToString();
    
    foreach (var field in new[] { "Father", "Mother", "Brother", "Sister" })
    {
        var val = values[field];
        if (!string.IsNullOrEmpty(val) && val != "null")
        {
            var parts = val.Split(',');
            person["RelativeFirstName"] = parts[0];
            person["RelativeLastName"] = parts[1];
        }
    }
    
    if (count > 0) 
        person.Remove("RelativeFirstName");
}

var jsonOutput = new List<Dictionary<string, object>>();
for (int i = 0; i < totalRows; i++)
{
    var values = lines[i + 1].Trim().Split(',');
    var firstName = values[0];
    var birthdayParts = values[1].Split('/');
    var birthYear = int.Parse(birthdayParts[2]);
    
    var age = currentYear - birthYear;
    if (age < 0) age = 0;
    
    var person = new Dictionary<string, object> { 
        {"FirstName", firstName},
        {"Birthday", $"{birthdayParts[0]}/{birthdayParts[1]}/{birthdayYear}"},
        {"Age", age}
    };
    
    foreach (var field in new[] { "Father", "Mother", "Brother", "Sister" })
    {
        var val = values[field];
        if (!string.IsNullOrEmpty(val) && val != "null")
        {
            var parts = val.Split(',');
            person["Relatives"] = ((List<object>)person["Relatives"]).Add(new Dictionary<string, object> { 
                {"FirstName", parts[0]}, 
                {"LastName", parts[1]},
                {"Relationship", field}
            }).ToList();
        }
    }
    
    jsonOutput.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(jsonOutput));