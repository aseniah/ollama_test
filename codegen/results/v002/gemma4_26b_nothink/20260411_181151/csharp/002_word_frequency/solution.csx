using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

// Check if file exists
string filePath = "input/text.txt";
if (!File.Exists(filePath))
{
    return;
}

// Read the entire file content
string content = File.ReadAllText(filePath);

// Initialize dictionary to store word frequencies
var wordCounts = new Dictionary<string, int>();

// Split by whitespace to get initial chunks
string[] chunks = content.Split(new[] { ' ', '\r', '\n', '\t' }, StringSplitOptions.RemoveEmptyEntries);

foreach (string chunk in chunks)
{
    // Remove all non-letter characters and convert to lowercase
    StringBuilder sb = new StringBuilder();
    foreach (char c in chunk)
    {
        if (char.IsLetter(c))
        {
            sb.Append(char.ToLowerInvariant(c));
        }
    }

    string cleanedWord = sb.ToString();

    // Only process if the word isn't empty after stripping punctuation
    if (!string.IsNullOrEmpty(cleanedWord))
    {
        if (wordCounts.ContainsKey(cleanedWord))
        {
            wordCounts[cleanedWord]++;
        }
        else
        {
            wordCounts[cleanedWord] = 1;
        }
    }
}

// Sorting logic:
// 1. Sort by count descending
// 2. Sort by word alphabetically ascending (for ties)
var sortedWords = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

// Output the results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}