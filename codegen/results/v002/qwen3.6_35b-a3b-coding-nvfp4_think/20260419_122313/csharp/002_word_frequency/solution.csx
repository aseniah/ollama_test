using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var lines = File.ReadAllLines("input/text.txt");
var counts = new Dictionary<string, int>();

foreach (var line in lines)
{
    foreach (Match match in Regex.Matches(line, @"[a-zA-Z]+"))
    {
        var word = match.Value.ToLower();
        counts.TryGetValue(word, out int count);
        counts[word] = count + 1;
    }
}

var sorted = counts.OrderByDescending(kvp => kvp.Value)
                   .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}