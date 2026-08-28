using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

// Define the file path
string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

// Read all content from the file
string content = File.ReadAllText(filePath);

// To strip punctuation and keep only letters, we can use a Regex or iterate through characters.
// We convert to lowercase first.
string lowercaseContent = content.ToLower();

// Replace any non-letter character with a space to ensure words are split correctly
StringBuilder sb = new StringBuilder();
foreach (char c in lowercaseContent)
{
    if (char.IsLetter(c))
    {
        sb.Append(c);
    }
    else
    {
        sb.Append(' ');
    }
}

string cleanContent = sb.ToString();

// Split by whitespace and remove empty entries
string[] words = cleanContent.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequency of each word
Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string word in words)
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

// Sort by count descending, then by word ascending
var sortedWords = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

// Output the results
foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}