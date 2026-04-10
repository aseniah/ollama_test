using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

// Read file content
string content = File.ReadAllText("input/text.txt");

// Extract words (alphanumeric only, keep only letters), convert to lowercase
var words = Regex.Matches(content, @"[\p{L}]+", MatchInfo).Select(m => m.Value.ToLowerInvariant()).ToList();

// Count frequencies
var wordCounts = words.GroupBy(w => w).ToDictionary(g => g.Key, g => g.Count());

// Convert to list for sorting
var sortedWords = wordCounts
    .Select(pair => new { Word = pair.Key, Count = pair.Value })
    .OrderByDescending(p => p.Count)
    .ThenBy(p => p.Word)
    .ToList();

// Output results
foreach (var item in sortedWords)
{
    Console.WriteLine($"{item.Word}:{item.Count}");
}