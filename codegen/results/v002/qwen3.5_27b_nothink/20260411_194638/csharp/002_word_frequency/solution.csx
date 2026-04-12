#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string path = "input/text.txt";

if (!File.Exists(path))
{
    throw new FileNotFoundException($"File not found: {path}");
}

string content = File.ReadAllText(path);

// Normalize to lowercase
content = content.ToLowerInvariant();

// Extract words: keep only letters, treat anything else as separator
string cleaned = string.Concat(content
    .Select(c => char.IsLetter(c) ? c : ' ')
);

string[] words = cleaned.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
var frequencies = new Dictionary<string, int>();

foreach (var word in words)
{
    if (string.IsNullOrEmpty(word)) continue;
    
    if (frequencies.ContainsKey(word))
    {
        frequencies[word]++;
    }
    else
    {
        frequencies[word] = 1;
    }
}

// Sort: Count descending, then Alphabetically ascending
var sorted = frequencies
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

// Output
foreach (var pair in sorted)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}