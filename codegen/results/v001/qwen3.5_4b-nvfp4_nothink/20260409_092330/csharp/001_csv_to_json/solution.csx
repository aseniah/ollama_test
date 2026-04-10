using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/data.csv";
if (File.Exists(inputPath) == false) {
    throw new FileNotFoundException("Input file not found.");
}

var content = File.ReadAllText(inputPath);
var lines = content.TrimEnd('\n').Split(new[] { '\r\n', '\n' }, StringSplitOptions.RemoveEmptyEntries);

if (lines.Length < 2) {
    Console.WriteLine('[]');
    return;
}

var array = JsonArray.Parse(@"[");

for (var i = 1; i < lines.Length; i++) {
    var parts = lines[i].Split(',', '\t');
    if (parts.Length >= 4) {
        try {
            var nameObj = Node.Create(parts[0]);
            var ageObj = Node.Create(parts[1])?.GetValueOrDefault(int.Parse)) ?? Node.Create("0");
            var emailObj = Node.Create(parts[2]);
            var scoreObj = Node.Create(parts[3])?.GetValueOrDefault(float.Parse)) ?? Node.Create(0.0f);

            array.Add(Node.Create(@"{""Name"": @nameObj, ""Age"": @ageObj, ""Email"": @emailObj, ""Score"": @scoreObj}"));
        } catch (Exception e) {
            // ignore parse errors for this script's simplicity
            array.Add(Node.Create(@"{""Name"": @parts[0], ""Age"": 0, ""Email"": @parts[2], ""Score"": 0.0}"));
        }
    } else {
        break;
    }
}

Console.WriteLine(array.ToString());