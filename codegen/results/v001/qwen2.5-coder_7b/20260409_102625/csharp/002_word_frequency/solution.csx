using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    Console.WriteLine("File not found.");
    return;
}

var text = File.ReadAllText(filePath);
var words = Regex.Matches(text.ToLower(), @"\b\w+\b")
                  .Cast<Match>()
                  .Select(m => m.Value)
                  .ToList();

var wordCounts = words.GroupBy(word => word)
                      .ToDictionary(group => group.Key, group => group.Count());

var sortedWordCounts = wordCounts.OrderByDescending(pair => pair.Value)
                                  .ThenBy(pair => pair.Key);

foreach (var pair in sortedWordCounts)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}