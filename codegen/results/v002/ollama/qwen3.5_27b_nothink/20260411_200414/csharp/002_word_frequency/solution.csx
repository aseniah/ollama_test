#r "System.Text.Json"

using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.RegularExpressions;

// Read the input file
string filePath = "input/text.txt";
string content = File.ReadAllText(filePath);

// Normalize text: lowercase
string lowerContent = content.ToLower();

// Split into potential words by whitespace, then filter out punctuation
// We keep only alphabetical characters in each segment
IEnumerable<string> rawTokens = lowerContent.Split((char[])null, StringSplitOptions.RemoveEmptyEntries);

// Create a dictionary to store word frequencies
Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string token in rawTokens)
{
    // Strip non-letters (keep only a-z)
    string cleanedWord = Regex.Replace(token, "[^a-z]", "");
    
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

// Sort: by count descending, then by word alphabetically ascending
var sortedWords = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

// Output results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}