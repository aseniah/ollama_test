using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath);

// Regular expression to keep only letters and whitespace, effectively stripping punctuation
// then split by whitespace to get words.
string cleanedText = Regex.Replace(text.ToLower(), @"[^a-z\s]", "");
string[] words = cleanedText.Split(new[] { ' ', '\t', '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> frequencies = new Dictionary<string, int>();

foreach (var word in words)
{
    if (frequencies.ContainsKey(word))
    {
        frequencies[word]++;
    }
    else
    {
        frequencies[word] = 1;
    }
}

// Sort by count descending, then by word alphabetically ascending
var sortedWords = frequencies
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}