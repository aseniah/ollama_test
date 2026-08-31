using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

string text = File.ReadAllText("input/text.txt");

// Convert to lowercase, strip punctuation, split into words
string lowerText = text.ToLower();

// We need to split on non-letter characters to get words
// First, replace all non-letter characters with spaces, then split
var words = lowerText
    .Select(c => char.IsLetter(c) ? c : ' ')
    .ToArray();

// Build words from the character array
var wordList = new List<string>();
var currentWord = new System.Text.StringBuilder();
foreach (char c in words)
{
    if (c == ' ')
    {
        if (currentWord.Length > 0)
        {
            wordList.Add(currentWord.ToString());
            currentWord.Clear();
        }
    }
    else
    {
        currentWord.Append(c);
    }
}
if (currentWord.Length > 0)
{
    wordList.Add(currentWord.ToString());
}

// Count frequencies
var counts = new Dictionary<string, int>();
foreach (var w in wordList)
{
    if (!counts.ContainsKey(w))
        counts[w] = 0;
    counts[w]++;
}

// Sort by count descending, then alphabetically ascending
var sorted = counts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}