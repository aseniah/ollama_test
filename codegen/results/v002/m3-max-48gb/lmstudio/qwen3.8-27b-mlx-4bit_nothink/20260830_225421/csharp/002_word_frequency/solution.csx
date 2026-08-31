using System;
using System.IO;
using System.Text;
using System.Linq;
using System.Collections.Generic;

var text = File.ReadAllText("input/text.txt");

// Convert to lowercase and strip all non-letter characters
var words = new List<string>();
var currentWord = new StringBuilder();

foreach (char c in text.ToLower())
{
    if (char.IsLetter(c))
    {
        currentWord.Append(c);
    }
    else
    {
        if (currentWord.Length > 0)
        {
            words.Add(currentWord.ToString());
            currentWord.Clear();
        }
    }
}
if (currentWord.Length > 0)
{
    words.Add(currentWord.ToString());
}

// Count frequencies
var counts = new Dictionary<string, int>();
foreach (var w in words)
{
    if (counts.ContainsKey(w))
        counts[w]++;
    else
        counts[w] = 1;
}

// Sort by count descending, then alphabetically ascending
var sorted = counts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}