using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

// Path to the input file
string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

// Read all text from the file
string content = File.ReadAllText(filePath);

// Use a StringBuilder to build a cleaned version of the string
// We replace everything that is not a letter or whitespace with a space
StringBuilder sb = new StringBuilder();
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

// Split the cleaned content into words by whitespace
string[] words = sb.ToString().Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count the frequency of each word
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

// Sort the dictionary: 
// 1. By count descending
// 2. By word ascending (alphabetical)
var sortedWords = counts.OrderByDescending(kvp => kvp.Value)
                        .ThenBy(kvp => kvp.Key);

// Output the results in the format "word: count"
foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}