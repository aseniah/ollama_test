using System;
using System.Globalization;
using System.IO;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++) {
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length >= 4) {
        var obj = new JsonObject();
        obj.Add("Name", parts[0].Trim());
        obj.Add("Age", int.Parse(parts[1].Trim()));
        obj.Add("Email", parts[2].Trim());
        obj.Add("Score", float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture));
        jsonArray.Add(obj);
    }
}

Console.WriteLine(jsonArray.ToJsonString());