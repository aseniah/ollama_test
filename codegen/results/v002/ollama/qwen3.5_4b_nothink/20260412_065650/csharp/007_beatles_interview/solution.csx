using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvLines = File.ReadAllLines("input/input.csv");
var lines = csvLines.Skip(1); // Skip header

var people = new List<object>();
var dateCulture = new CultureInfo("en-US");

foreach (var line in lines) {
    var parts = line.Split(',');
    var name = parts[0];
    var birthdayStr = parts[1];
    var died = parts[2];
    var father = parts[3];
    var mother = parts[4];
    var brother = parts[5];
    var sister = parts[6];

    // Parse birthday
    var birthdayParts = birthdayStr.Split('/');
    var birthDate = new DateTime(int.Parse(birthdayParts[2], dateCulture), int.Parse(birthdayParts[1], dateCulture), int.Parse(birthdayParts[0], dateCulture));

    // Calculate age as of July 1, 2025
    var targetDate = new DateTime(2025, 7, 1);
    var age = (int)(targetDate - birthDate).TotalDays;

    // Parse relatives
    var relatives = new List<object>();
    if (!string.IsNullOrEmpty(father)) {
        var p = new { FirstName = father, Relationship = "Father" };
        relatives.Add(p);
    }
    if (!string.IsNullOrEmpty(mother)) {
        var p = new { FirstName = mother, Relationship = "Mother" };
        relatives.Add(p);
    }
    if (!string.IsNullOrEmpty(brother)) {
        var p = new { FirstName = brother, Relationship = "Brother" };
        relatives.Add(p);
    }
    if (!string.IsNullOrEmpty(sister)) {
        var p = new { FirstName = sister, Relationship = "Sister" };
        relatives.Add(p);
    }

    var person = new {
        FirstName = name.Split(' ')[0],
        LastName = name.Split(' ').Last(),
        Birthday = $"{birthDate:yyyy}-{birthDate:MM}-{birthDate:dd}",
        Age = age,
        Relatives = relatives
    };

    people.Add(person);
}

var json = JsonNode.CreateArray(people);
Console.WriteLine(JsonSerializer.Serialize(json, new JsonSerializerOptions { WriteIndented = true }));