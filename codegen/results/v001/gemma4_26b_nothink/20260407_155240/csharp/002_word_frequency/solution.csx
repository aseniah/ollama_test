using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Collections.Generic;

// Define the path to the input file
string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    // If the file doesn't exist, we exit silently as per the requirement to output nothing else
    return;
}

// Read the entire content of the file
string content = File.ReadAllText(filePath);

// Use a StringBuilder for efficient processing
StringBuilder sb = new StringBuilder();

// Process the text: keep only letters and spaces
// We replace everything that isn't a letter or a space with a space to avoid merging words
foreach (char c in content)
{
    if (char.IsLetter(c) || char.IsWhiteSpace(c))
    {
        sb.Append(char.ToLowerInvariant(c));
    }
    else
    {
        sb.Append(' ');
    }
}

// Split by whitespace and remove empty entries
string[] words = sb.ToString().Split(new[] { ' ', '\r', '\n', '\t' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies using a Dictionary
Dictionary<string, int> counts = new Dictionary<string, int>();

foreach (string word in words)
{
    if (counts.ContainsKey(word))
    {
        counts[word]++;
    }
    else
    {
        counts[word] = 1;
    }
}

// Sort the results:
// 1. By count descending
// 2. By word alphabetically ascending
var sortedWords = counts
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

// Output the formatted lines
foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}