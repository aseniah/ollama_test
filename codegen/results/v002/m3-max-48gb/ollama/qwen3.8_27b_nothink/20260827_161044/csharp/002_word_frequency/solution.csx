using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

string[] lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (string line in lines)
{
    // Split by non-letter characters
    string[] words = System.Text.RegularExpressions.Regex.Split(line, @"[^a-zA-Z]+");
    foreach (string w in words)
    {
        if (string.IsNullOrEmpty(w)) continue;
        string word = w.ToLower();
        if (wordCounts.ContainsKey(word))
            wordCounts[word]++;
        else
            wordCounts[word] = 1;
    }
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts.OrderByDescending(kv => kv.Value)
                       .ThenBy(kv => kv.Key)
                       .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}