using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var results = new List<JsonObject>();

var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    var person = new JsonObject();
    
    // Parse name
    var fullName = fields[0].Trim();
    var nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    person["FirstName"] = nameParts[0];
    person["LastName"] = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
    // Parse birthday
    var birthDate = DateTime.Parse(fields[1].Trim());
    person["Birthday"] = birthDate.ToString("yyyy-MM-dd");
    
    // Calculate age as of July 1, 2025
    var age = referenceDate.Year - birthDate.Year;
    if (referenceDate.Month < birthDate.Month || 
        (referenceDate.Month == birthDate.Month && referenceDate.Day < birthDate.Day))
    {
        age--;
    }
    person["Age"] = age;
    
    // Parse relatives
    var relatives = new JsonArray();
    
    // Father
    if (!string.IsNullOrEmpty(fields[3].Trim()))
    {
        var father = new JsonObject();
        var fatherName = fields[3].Trim();
        var fatherNameParts = fatherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        father["FirstName"] = fatherNameParts[0];
        father["LastName"] = fatherNameParts.Length > 1 ? string.Join(" ", fatherNameParts.Skip(1)) : "";
        father["Relationship"] = "Father";
        relatives.Add(father);
    }
    
    // Mother
    if (!string.IsNullOrEmpty(fields[4].Trim()))
    {
        var mother = new JsonObject();
        var motherName = fields[4].Trim();
        var motherNameParts = motherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        mother["FirstName"] = motherNameParts[0];
        mother["LastName"] = motherNameParts.Length > 1 ? string.Join(" ", motherNameParts.Skip(1)) : "";
        mother["Relationship"] = "Mother";
        relatives.Add(mother);
    }
    
    // Brothers
    if (!string.IsNullOrEmpty(fields[5].Trim()))
    {
        var brotherName = fields[5].Trim();
        if (brotherName != "null")
        {
            var brother = new JsonObject();
            var brotherNameParts = brotherName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            brother["FirstName"] = brotherNameParts[0];
            brother["LastName"] = brotherNameParts.Length > 1 ? string.Join(" ", brotherNameParts.Skip(1)) : "";
            brother["Relationship"] = "Brother";
            relatives.Add(brother);
        }
    }
    
    // Sisters
    if (!string.IsNullOrEmpty(fields[6].Trim()))
    {
        var sisterName = fields[6].Trim();
        if (sisterName != "null")
        {
            var sister = new JsonObject();
            var sisterNameParts = sisterName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            sister["FirstName"] = sisterNameParts[0];
            sister["LastName"] = sisterNameParts.Length > 1 ? string.Join(" ", sisterNameParts.Skip(1)) : "";
            sister["Relationship"] = "Sister";
            relatives.Add(sister);
        }
    }
    
    person["Relatives"] = relatives;
    results.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(results, new JsonSerializerOptions { WriteIndented = true }));