#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string content = File.ReadAllText("input/text.txt");

// Normalize: convert to lowercase and keep only letters, replacing others with spaces
StringBuilder sb = new StringBuilder();
foreach (char c in content)
{
    if (char.IsLetter(c))
    {
        sb.Append(char.ToLower(c));
    }
    else
    {
        sb.Append(' ');
    }
}

string normalized = sb.ToString();
string[] words = normalized.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
var counts = new Dictionary<string, int>();

foreach (string word in words)
{
    if (word.Length == 0) continue;
    
    if (counts.ContainsKey(word))
    {
        counts[word]++;
    }
    else
    {
        counts[word] = 1;
    }
}

// Sort: by count descending, then alphabetically ascending
var sortedWords = counts.OrderByDescending(k => k.Value)
                        .ThenBy(k => k.Key);

// Output
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}