using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string content = File.ReadAllText(filePath);

// Split by whitespace to get raw tokens
string[] rawWords = content.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> frequencies = new Dictionary<string, int>();

foreach (string rawWord in rawWords)
{
    // Convert to lowercase and remove non-letter characters
    string cleaned = Regex.Replace(rawWord.ToLower(), "[^a-z]", "");
    
    if (!string.IsNullOrEmpty(cleaned))
    {
        if (frequencies.ContainsKey(cleaned))
        {
            frequencies[cleaned]++;
        }
        else
        {
            frequencies[cleaned] = 1;
        }
    }
}

// Sort by count descending, then by word ascending
var sortedWords = frequencies
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}