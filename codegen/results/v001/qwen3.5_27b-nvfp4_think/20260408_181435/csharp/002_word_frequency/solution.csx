using System;
using System.IO;
using System.Linq;

var text = File.ReadAllText("input/text.txt");

// Split by whitespace, convert to lowercase, strip non-letters, filter empty
var words = text
    .Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => w.ToLower())
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()))
    .Where(w => w.Length > 0);

// Count frequencies
var wordCounts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

// Sort by count descending, then alphabetically ascending for ties
var sortedWords = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key);

// Output
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}