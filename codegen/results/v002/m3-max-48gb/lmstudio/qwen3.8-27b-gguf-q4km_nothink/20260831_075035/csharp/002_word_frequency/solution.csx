using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string text = File.ReadAllText("input/text.txt");

// Split into words by whitespace
var words = text
    .Split(new[] { ' ', '\t', '\n', '\r', '\f', '\v' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()))
    .Where(w => w.Length > 0)
    .Select(w => w.ToLower());

var counts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sorted = counts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}