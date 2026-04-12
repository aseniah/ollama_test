#r "System.Text.RegularExpressions"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

// Read the file
string content = File.ReadAllText("input/text.txt");

// Convert to lowercase
content = content.ToLower();

// Keep only letters (a-z) - replace everything else with space
content = Regex.Replace(content, "[^a-z]", " ");

// Split into words (handles multiple spaces)
string[] words = content.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
Dictionary<string, int> wordCount = new Dictionary<string, int>();
foreach (string word in words)
{
    if (word.Length > 0)
    {
        if (wordCount.ContainsKey(word))
        {
            wordCount[word]++;
        }
        else
        {
            wordCount[word] = 1;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCount.OrderByDescending(x => x.Value)
                      .ThenBy(x => x.Key)
                      .ToList();

// Output
foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}