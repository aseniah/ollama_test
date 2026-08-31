using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;

string text = File.ReadAllText("input/text.txt");
string[] rawWords = text.Split(new[] { ' ', '\t', '\n', '\r', ',' }, StringSplitOptions.RemoveEmptyEntries);

var wordCounts = new Dictionary<string, int>(StringComparer.OrdinalIgnoreCase);

foreach (string rawWord in rawWords)
{
    // Keep only letters, convert to lowercase
    StringBuilder sb = new StringBuilder();
    foreach (char c in rawWord)
    {
        if (char.IsLetter(c))
            sb.Append(char.ToLower(c));
    }
    string word = sb.ToString();
    if (word.Length > 0)
    {
        if (wordCounts.ContainsKey(word))
            wordCounts[word]++;
        else
            wordCounts[word] = 1;
    }
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}