using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

if (!File.Exists("input/text.txt"))
{
    Console.WriteLine("Error: input/text.txt not found.");
    return;
}

string text = File.ReadAllText("input/text.txt");

// Tokenize: Split by whitespace and other non-letter characters
// Using regex-like approach by extracting sequences of letters
var words = System.Text.RegularExpressions.Regex.Matches(text, "[a-zA-Z]+")
             .Cast<System.Text.RegularExpressions.Match>()
             .Select(m => m.Value.ToLower())
             .Where(w => w.Length > 0);

var counts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

// Sort by count descending, then word alphabetically ascending
var sorted = counts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}