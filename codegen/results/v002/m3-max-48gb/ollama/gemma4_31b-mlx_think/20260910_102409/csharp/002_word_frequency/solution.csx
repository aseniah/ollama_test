using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath);
// Split by any whitespace character
string[] rawWords = text.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string rawWord in rawWords)
{
    // Strip punctuation: keep only letters and convert to lowercase
    StringBuilder sb = new StringBuilder();
    foreach (char c in rawWord)
    {
        if (char.IsLetter(c))
        {
            sb.Append(char.ToLowerInvariant(c));
        }
    }

    string word = sb.ToString();
    if (!string.IsNullOrEmpty(word))
    {
        if (wordCounts.ContainsKey(word))
        {
            wordCounts[word]++;
        }
        else
        {
            wordCounts[word] = 1;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sortedWords = wordCounts
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}