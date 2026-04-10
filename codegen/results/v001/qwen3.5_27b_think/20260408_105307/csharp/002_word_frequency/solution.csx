#r "System.Text.RegularExpressions"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var content = File.ReadAllText("input/text.txt");

var words = Regex.Matches(content.ToLower(), @"[a-z]+")
    .Cast<Match>()
    .Select(m => m.Value);

var counts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sorted = counts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key);

foreach (var item in sorted)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}