using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

var lines = File.ReadAllLines("input/data.csv");
var results = new List<string>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var fields = line.Split(',');
    var name = fields[0];
    var age = int.Parse(fields[1]);
    var email = fields[2];
    var score = float.Parse(fields[3]);
    
    var json = $"{{\"Name\":\"{name}\",\"Age\":{age},\"Email\":\"{email}\",\"Score\":{score.ToString("F1", System.Globalization.CultureInfo.InvariantCulture)}}}";
    results.Add(json);
}

var output = "[" + string.Join(",", results) + "]";
Console.Write(output);