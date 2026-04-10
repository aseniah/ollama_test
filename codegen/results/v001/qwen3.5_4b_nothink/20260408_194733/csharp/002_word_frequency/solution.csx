using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var lines = File.ReadAllLines("input/text.txt");
string content = string.Join("", lines);

// Convert to lowercase
content = content.ToLowerInvariant();

// Extract words: sequences of letters (keep only letters)
var matches = Regex.Matches(content, "[a-z]+");

// Count frequencies
var wordCounts = new Dictionary<string, int>();
foreach (Match match in matches)
{
    string word = match.Value;
    if (!wordCounts.ContainsKey(word))
    {
        wordCounts[word] = 0;
    }
    wordCounts[word]++;
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts
    .ToDictionary(k => k.Key, v => v.Value,
        (x, y) => x.Value.CompareTo(y.Value) != 0
            ? x.Value.CompareTo(y.Value)
            : x.Key.CompareTo(y.Key))
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

// Output
foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}