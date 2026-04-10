using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

List<string> lines = File.ReadAllLines("input/data.csv").ToList();
if (lines.Count <= 0) {
    Console.WriteLine("[]");
    return;
}

string[] header = lines[0].Split(',);
List<Dictionary<string, string>> rows = new List<Dictionary<string, string>>();
for (int i = 1; i < lines.Count; i++) {
    string[] row = lines[i].Split(',);
    Dictionary<string, string> rowDict = new Dictionary<string, string>();
    for (int j = 0; j < row.Length; j++) {
        string name = header[j].Trim();
        string value = row[j].Trim();
        rowDict[name] = value;
    }
    rows.Add(rowDict);
}

List<JsonElement> results = new List<JsonElement>();
foreach (var row in rows) {
    string name = row["Name"];
    int age = int.Parse(row["Age"], CultureInfo.InvariantCulture);
    string email = row["Email"];
    float score = float.Parse(row["Score"], CultureInfo.InvariantCulture);
    JsonElement obj = JsonNode.Parse($"{{"name": "{name}","age": {age},"email": "{email}","score": {score}}}".Replace("{0}", "0"));
    results.Add(obj);
}

Console.WriteLine(JsonConvert.ToString(results));