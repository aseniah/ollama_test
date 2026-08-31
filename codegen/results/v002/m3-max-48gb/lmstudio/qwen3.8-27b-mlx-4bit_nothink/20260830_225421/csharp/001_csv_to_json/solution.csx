using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length < 4) continue;
    
    var name = parts[0];
    var age = int.Parse(parts[1], CultureInfo.InvariantCulture);
    var email = parts[2];
    var score = float.Parse(parts[3], CultureInfo.InvariantCulture);
    
    var obj = new JsonObject
    {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score
    };
    
    jsonArray.Add(obj);
}

var options = new JsonWriterOptions
{
    Indented = false
};

using var writer = new Utf8JsonWriter(new MemoryStream());
writer.WriteStartObject();
writer.WritePropertyName("dummy");
writer.WriteEndObject();

// Actually, let's just use the JsonArray's ToJsonString
Console.Write(jsonArray.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));