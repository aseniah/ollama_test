using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var filePath = "input/text.txt";
var text = File.ReadAllText(filePath);

var wordCounts = new Dictionary<string, int>();

var words = Regex.Matches(text, @"[a-z]+", RegexOptions.IgnoreCase)
    .Cast<Match>()
    .Select(m => m.Value.ToLower());

foreach (var word in words)
{
    if (wordCounts.ContainsKey(word))
        wordCounts[word]++;
    else
        wordCounts[word] = 1;
}

var sorted = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}
