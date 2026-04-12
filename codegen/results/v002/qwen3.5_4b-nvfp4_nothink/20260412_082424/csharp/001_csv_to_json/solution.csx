using System;
using System.Text.Json.Nodes;
using System.Text.Json;
using System.Collections.Generic;
using System.IO;

var lines = File.ReadAllLines("input/data.csv");
int headerIndex = 0;
for (int i = 0; i < lines.Length; i++) {
    if (string.IsNullOrEmpty(lines[i])) continue;
    headerIndex = i;
    break;
}
var firstRow = lines[headerIndex];
if (string.IsNullOrEmpty(firstRow)) continue;

var firstParts = firstRow.Split(',');
int numFields = firstParts.Length;

var output = new List<object>();
for (int i = headerIndex + 1; i < lines.Length; i++) {
    if (string.IsNullOrEmpty(lines[i])) continue;
    var parts = lines[i].Split(',');
    if (parts.Length != numFields) continue;

    string name = parts[0];
    int age = 0;
    if (int.TryParse(parts[1], out int tempAge)) {
        age = tempAge;
    } else {
        continue;
    }
    
    string email = parts[2];
    float score = 0.0f;
    if (float.TryParse(parts[3], out float tempScore)) {
        score = tempScore;
    } else {
        continue;
    }

    var rowObj = new JsonObject();
    rowObj["Name"] = name;
    rowObj["Age"] = age;
    rowObj["Email"] = email;
    rowObj["Score"] = score;
    
    output.Add(rowObj);
}

var jsonDocument = new JsonDocument();
jsonDocument.ParseArray(new object[1] { null }); // placeholder to start
// Re-build properly using array
jsonDocument = new JsonDocument(JsonSerializer.Serialize(output, FormatOption.PreserveReferencing)));
// Actually need to serialize each row correctly
var jsonString = JsonSerializer.Serialize(output);
Console.WriteLine(jsonString);