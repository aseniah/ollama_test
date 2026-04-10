using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

var content = File.ReadAllText("input/text.txt");

var words = Regex.Matches(content.ToLowerInvariant(), @"[\p{L}]+")
    .Select(m => m.Value)
    .ToArray();

var counts = new Dictionary<string, int>();
foreach (var word in words) {
    if (counts.ContainsKey(word)) {
        counts[word]++;
    } else {
        counts.Add(word, 1);
    }
}

var result = counts
    .Select(kvp => new { Word = kvp.Key, Count = kvp.Value })
    .OrderByDescending(p => p.Count)
    .ThenBy(p => p.Word)
    .ToList();

foreach (var item in result) {
    Console.WriteLine($"{item.Word}: {item.Count}");
}