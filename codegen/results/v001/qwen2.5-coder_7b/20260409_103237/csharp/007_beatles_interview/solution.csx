using System;
using System.IO;
using System.Linq;
using System.Text.Json;

var birthdates = new[] { "1990-05-15", "1985-12-20", "2000-07-01" }; // Example birthdates
var currentDate = DateTime.Parse("2025-07-01");

var jsonNodes = birthdates.Select(birthdate => {
    var birthDate = DateTime.Parse(birthdate);
    var age = currentDate.Year - birthDate.Year;
    if (currentDate < birthDate.AddYears(age)) age--;
    return JsonNode.Parse($"{{\"birthdate\": \"{birthdate}\", \"age\": {age}}}");
}).ToArray();

var jsonArray = new JsonArray(jsonNodes);
Console.WriteLine(jsonArray);