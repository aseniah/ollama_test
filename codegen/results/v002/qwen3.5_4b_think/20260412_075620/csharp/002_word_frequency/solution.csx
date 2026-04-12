using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

var text = File.ReadAllText("input/text.txt");
var counts = new Dictionary<string, int>();

foreach (var line in text.Split(new[] { '\n', '\r' }, StringSplitOptions.None)) {
    foreach (var word in line.Split()) {
        string w = word.ToLowerInvariant();
        // Strip all punctuation (keep only letters)
        string clean = w.Where(char.IsLetter).ToArray().ToString();
        if (!string.IsNullOrEmpty(clean)) {
            if (counts.ContainsKey(clean)) {
                counts[clean]++;
            } else {
                counts[clean] = 1;
            }
        }
    }
}

var sorted = counts.Select(kvp => new { Word = kvp.Key, Count = kvp.Value })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word)
    .ToList();

foreach (var item in sorted) {
    Console.WriteLine($"{item.Word}: {item.Count}");
}