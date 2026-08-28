using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string content = File.ReadAllText("input/text.txt");
string[] words = content.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

var counts = new Dictionary<string, int>();

foreach (string raw in words)
{
    // Strip all non-letter characters, convert to lowercase
    string word = new string(raw.Where(c => char.IsLetter(c)).ToArray()).ToLower();
    if (string.IsNullOrEmpty(word)) continue;
    
    if (counts.ContainsKey(word))
        counts[word]++;
    else
        counts[word] = 1;
}

var sorted = counts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}