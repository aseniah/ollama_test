using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

if (!File.Exists("input/text.txt"))
{
    return;
}

string content = File.ReadAllText("input/text.txt");

// Split into words by whitespace
string[] rawWords = content.Split(new[] { ' ', '\r', '\n', '\t' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string rawWord in rawWords)
{
    // Convert to lowercase and strip non-letters
    string cleaned = Regex.Replace(rawWord.ToLower(), @"[^a-z]", "");

    if (!string.IsNullOrEmpty(cleaned))
    {
        if (wordCounts.ContainsKey(cleaned))
        {
            wordCounts[cleaned]++;
        }
        else
        {
            wordCounts[cleaned] = 1;
        }
    }
}

// Sort by count descending, then by word ascending
var sortedWords = wordCounts
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}