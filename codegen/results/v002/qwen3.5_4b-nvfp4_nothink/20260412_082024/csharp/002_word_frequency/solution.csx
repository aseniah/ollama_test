using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;

var lines = File.ReadAllLines("input/text.txt");
var words = string.Join("\n", lines)
    .ToLower()
    .Replace("[^a-z]", "", RegexOptions.None) // Keep only letters, convert to lowercase
    .Split(' ')
    .Where(w => !string.IsNullOrEmpty(w)) // Filter empty strings
    .Select(w => w.ToLower()) // Double-safeguard lowercase
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sorted = words
    .OrderByDescending(w => w.Value)
    .ThenBy(w => w.Key)
    .ToList();

foreach (var word in sorted)
{
    Console.WriteLine($"{word.Key}: {word.Value}");
}