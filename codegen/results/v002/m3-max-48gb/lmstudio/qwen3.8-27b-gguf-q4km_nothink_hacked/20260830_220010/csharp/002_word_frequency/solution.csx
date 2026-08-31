using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;

string content = File.ReadAllText("input/text.txt");

// Keep only letters and spaces (strip punctuation, convert to lowercase)
string cleaned = string.Concat(content.Where(c => char.IsLetter(c) || char.IsWhiteSpace(c)).Select(c => char.ToLowerInvariant(c)));

// Split on whitespace
string[] words = cleaned.Split(new char[] { ' ', '\t', '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequency
var freq = new Dictionary<string, int>();
foreach (string w in words)
{
    if (freq.ContainsKey(w))
        freq[w]++;
    else
        freq[w] = 1;
}

// Sort by count descending, then alphabetically ascending
var sorted = freq.OrderByDescending(kvp => kvp.Value)
                 .ThenBy(kvp => kvp.Key)
                 .ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}