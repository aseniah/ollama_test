using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var text = File.ReadAllText("input/text.txt");

// Split into words, strip punctuation (keep only letters), lowercase
var words = text
    .Split(new[] {' ', ',', '\n', '\r', '.', '!', '?'}, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()).ToLower())
    .Where(w => w.Length > 0);

var frequency = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

// Sort by count descending, then alphabetically ascending
var sorted = frequency
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key, StringComparer.Ordinal)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}