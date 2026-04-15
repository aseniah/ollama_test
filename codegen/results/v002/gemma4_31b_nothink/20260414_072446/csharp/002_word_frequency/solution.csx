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

// Normalize: lowercase
string lowerContent = content.ToLower();

// Strip punctuation: Keep only letters and whitespace
StringBuilder sb = new StringBuilder();
foreach (char c in lowerContent)
{
    if (char.IsLetter(c) || char.IsWhiteSpace(c))
    {
        sb.Append(c);
    }
    else
    {
        sb.Append(' '); // Replace punctuation with space to avoid merging words
    }
}

// Split into words and filter out empty entries
string[] words = sb.ToString()
    .Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
Dictionary<string, int> frequencies = new Dictionary<string, int>();
foreach (string word in words)
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

// Sort by count descending, then by word ascending
var sortedWords = frequencies
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}