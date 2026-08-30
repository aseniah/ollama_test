using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using System.Collections.Generic;

var content = File.ReadAllText("input/text.txt");
var words = Regex.Matches(content.ToLower(), @"[a-z]+").Cast<string>().ToArray();

var counts = words.GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count())
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var kvp in counts)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}