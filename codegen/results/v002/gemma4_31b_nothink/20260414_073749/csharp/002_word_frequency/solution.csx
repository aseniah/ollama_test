using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

if (!File.Exists("input/text.txt"))
{
    return;
}

string content = File.ReadAllText("input/text.txt");

// Use a regex to find sequences of alphabetic characters
// This effectively strips punctuation and splits by whitespace/non-letters
var matches = Regex.Matches(content.ToLower(), @"[a-z]+");

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (Match match in matches)
{
    string word = match.Value;
    if (wordCounts.ContainsKey(word))
    {
        wordCounts[word]++;
    }
    else
    {
        wordCounts[word] = 1;
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