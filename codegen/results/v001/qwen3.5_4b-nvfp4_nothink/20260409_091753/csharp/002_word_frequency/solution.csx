using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string inputPath = "input/text.txt";
string content = File.ReadAllText(inputPath);

var words = new List<string>();

foreach (char c in content) {
    string ch = c.ToString();
    if (!char.IsLetterOrDigit(c)) {
        continue;
    }
    if (char.IsUpper(ch)) {
        ch = char.ToLower(ch);
    }
    words.Add(ch);
}

var counts = words.ToDictionary(w => w, _ => 1).GroupBy(k => k.Key, _, _)
    .Select(g => new KeyValuePair<string, int>(g.Key, g.Count()))
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

var result = string.Join("\n", counts.Select(x => $"{x.Key}: {x.Value}"));
string jsonOutput = "{\"output\": \"" + result.Replace("\"", "\\\"") + "\"}";
Console.WriteLine(jsonOutput);