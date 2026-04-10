using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

// Read file content
string content = File.ReadAllText("input/text.txt");

// Extract words (only letters, lowercase)
var words = Regex.Split(content.ToLower(), "[^a-z]+")
    .Where(w => !string.IsNullOrEmpty(w))
    .ToList();

// Count and sort
var wordCounts = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word);

// Output
foreach (var item in wordCounts)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}