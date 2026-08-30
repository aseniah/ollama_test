#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

// Read the input file
var lines = File.ReadAllLines("input/text.txt");
string content = string.Join(" ", lines);

// Normalize to lowercase
content = content.ToLowerInvariant();

// Split into words and strip punctuation (keep only letters)
var words = Regex.Matches(content, "[a-z]+")
                 .Cast<Match>()
                 .Select(m => m.Value);

// Count frequencies
var frequencyDict = new Dictionary<string, int>();

foreach (var word in words)
{
    if (!string.IsNullOrEmpty(word))
    {
        if (frequencyDict.ContainsKey(word))
        {
            frequencyDict[word]++;
        }
        else
        {
            frequencyDict[word] = 1;
        }
    }
}

// Sort: first by count descending, then by word alphabetically ascending
var sortedItems = frequencyDict
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

// Output the result
foreach (var item in sortedItems)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}