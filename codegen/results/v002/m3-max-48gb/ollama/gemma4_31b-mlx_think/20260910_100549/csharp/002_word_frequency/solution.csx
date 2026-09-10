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

string content = File.ReadAllText(filePath);
// Split by whitespace characters
string[] tokens = content.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string token in tokens)
{
    StringBuilder sb = new StringBuilder();
    foreach (char c in token)
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

// Sort by count descending, then by word ascending
var sortedWords = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}