using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json).AsArray();

var filteredArray = array
    .Select(obj => (name: obj["name"].GetString(), age: obj["age"].GetInt32(), active: obj["active"].GetBoolean(), score: obj["score"].GetDouble()))
    .Where(record => record.active && record.age >= 30)
    .OrderBy(record => record.name)
    .Select(record => new { record.name, record.age, record.active, record.score })
    .ToArray();

var resultJson = JsonSerializer.Serialize(filteredArray, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(resultJson);