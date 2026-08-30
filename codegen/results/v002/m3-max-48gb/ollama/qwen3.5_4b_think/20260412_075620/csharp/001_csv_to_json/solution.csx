using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");

var data = new List<Dictionary<string, object>>();
foreach (var line in lines.Skip(1)) {
    var parts = line.Split(',');
    data.Add(new Dictionary<string, object> {
        { "Name", parts[0] },
        { "Age", int.Parse(parts[1]) },
        { "Email", parts[2] },
        { "Score", float.Parse(parts[3]) }
    });
}

Console.WriteLine(JsonSerializer.Serialize(data));