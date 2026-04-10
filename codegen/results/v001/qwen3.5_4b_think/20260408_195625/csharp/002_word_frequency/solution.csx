using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string content = File.ReadAllText("input/text.txt");

// Split into words and filter
var words = new List<string>();
using (var reader = new StringReader(content))
{
    while (true)
    {
        string line = reader.ReadLine();
        if (string.IsNullOrEmpty(line)) break;
        var token = line.Trim().ToLower();
        token = Regex.Replace(token, @"[^a-z0-9]", "");
        if (!string.IsNullOrEmpty(token))
        {
            words.Add(token);
        }
    }
}

// Count frequencies
var wordCounts = new Dictionary<string, int>();
foreach (var word in words)
{
    if (!wordCounts.ContainsKey(word))
    {
        wordCounts[word] = 0;
    }
    wordCounts[word]++;
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

// Output results
foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}