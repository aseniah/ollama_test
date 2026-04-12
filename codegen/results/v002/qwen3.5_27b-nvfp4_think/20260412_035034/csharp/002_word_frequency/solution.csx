using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string content = File.ReadAllText("input/text.txt");

// Convert to lowercase and extract words (letters only)
var words = Regex.Matches(content.ToLower(), @"[a-z]+")
    .Cast<Match>()
    .Select(m => m.Value)
    .ToList();

// Count frequencies
var frequency = new Dictionary<string, int>();
foreach (var word in words)
{
    if (frequency.ContainsKey(word))
        frequency[word]++;
    else
        frequency[word] = 1;
}

// Sort by count descending, then alphabetically ascending for ties
var sorted = frequency
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

// Output results
foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}