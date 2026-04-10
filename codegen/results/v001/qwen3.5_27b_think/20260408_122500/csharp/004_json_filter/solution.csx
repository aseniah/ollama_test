using System;
using System.IO;
using System.Text.Json;
using System.Linq;
using System.Collections.Generic;

public class Person
{
    public string Name { get; set; }
    public int Age { get; set; }
    public bool Active { get; set; }
    public float Score { get; set; }
}

var jsonContent = File.ReadAllText("input/data.json");
var records = JsonSerializer.Deserialize<List<Person>>(jsonContent) ?? new List<Person>();

var filtered = records
    .Where(p => p.Active && p.Age >= 30)
    .OrderBy(p => p.Name)
    .ToList();

Console.Write(JsonSerializer.Serialize(filtered));