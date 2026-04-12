using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

var lines = File.ReadAllLines("input/text.txt");
var text = string.Join(" ", lines.Select(l => Regex.Split(l.ToLower(), @"\W+")));

var words = new List<string>();
foreach (var line in text) {
    var match = Regex.Matches(line);
    foreach (Match m in match) {
        words.Add(m.Value.Trim());
    }
}

var counts = new Dictionary<string, int>();
foreach (var word in words) {
    counts[word] = counts[word.GetValueOrDefault(0)] + 1;
}

var sorted = counts.OrderByDescending(kv => kv.Value)
                   .ThenBy(kv => kv.Key)
                   .ToList();

foreach (var kvp in sorted) {
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}