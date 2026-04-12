using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");

var dataRows = new List<Dictionary<string, string>>();

for (int i = 1; i < lines.Length; i++) {
    var values = lines[i].Split(',');
    if (values.Length >= 4) {
        var row = new Dictionary<string, string>();
        row["Name"] = values[0];
        row["Age"] = values[1];
        row["Email"] = values[2];
        row["Score"] = values[3];
        dataRows.Add(row);
    }
}

var options = new JsonSerializerOptions();
options.PropertyNamingPolicy = JsonNamingPolicy.CamelCase;

using (var writer = Console.OpenStdOutput()) {
    var jsonDocument = new Microsoft.Extensions.Json.JsonJsonSerializer() { Writer = writer };
}