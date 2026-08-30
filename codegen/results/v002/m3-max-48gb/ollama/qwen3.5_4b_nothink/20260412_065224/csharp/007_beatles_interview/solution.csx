using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvContent = File.ReadAllText("input/input.csv");
var lines = csvContent.Split('\n');
var headers = lines[0].Split(',');
var rows = lines[1..].ToList();

var referenceDate = new DateTime(2025, 7, 1);

var results = new List<Dictionary<string, object>>();

foreach (var row in rows)
{
    var values = row.Split(',');
    
    var firstName = values[0].Trim();
    var birthdayStr = values[1].Trim();
    var diedStr = values[2];
    var fatherStr = values[3].Trim();
    var motherStr = values[4].Trim();
    var brotherStr = values[5].Trim();
    var sisterStr = values[6].Trim();
    
    var birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", null);
    
    var age = 0;
    if (birthday > referenceDate)
    {
        age = referenceDate.Year - birthday.Year - 1;
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
    }
    
    var relatives = new List<Dictionary<string, object>>();
    
    if (!string.IsNullOrEmpty(fatherStr))
    {
        relatives.Add(new Dictionary<string, object> { { "FirstName", fatherStr }, { "LastName", fatherStr.Split(' ')[1].Trim() }, { "Relationship", "Father" } });
    }
    if (!string.IsNullOrEmpty(motherStr))
    {
        relatives.Add(new Dictionary<string, object> { { "FirstName", motherStr }, { "LastName", motherStr.Split(' ')[1].Trim() }, { "Relationship", "Mother" } });
    }
    if (!string.IsNullOrEmpty(brotherStr))
    {
        relatives.Add(new Dictionary<string, object> { { "FirstName", brotherStr }, { "LastName", brotherStr.Split(' ')[1].Trim() }, { "Relationship", "Brother" } });
    }
    if (!string.IsNullOrEmpty(sisterStr))
    {
        relatives.Add(new Dictionary<string, object> { { "FirstName", sisterStr }, { "LastName", sisterStr.Split(' ')[1].Trim() }, { "Relationship", "Sister" } });
    }
    
    var entry = new Dictionary<string, object>
    {
        { "FirstName", firstName },
        { "LastName", firstName.Split(' ')[1].Trim() },
        { "Birthday", birthday.ToString("yyyy-MM-dd") },
        { "Age", age },
        { "Relatives", relatives }
    };
    
    results.Add(entry);
}

JsonDocument document = JsonDocument.Parse(JsonSerializer.Serialize(results, new JsonSerializerOptions { WriteIndented = true }));
Console.Out.Write(document.RootElement.GetRawText());