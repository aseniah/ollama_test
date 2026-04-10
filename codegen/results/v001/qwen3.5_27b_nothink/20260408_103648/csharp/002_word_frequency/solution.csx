#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

string inputFile = "input/text.txt";

if (!File.Exists(inputFile))
{
    // In a script context, we might want to handle missing files gracefully or just let it fail.
    // Based on the prompt, we assume the file exists or throw.
    Console.Error.WriteLine($"Error: File '{inputFile}' not found.");
    Environment.Exit(1);
}

string content = File.ReadAllText(inputFile);

// Convert to lowercase
string lowerContent = content.ToLower();

// Strip all punctuation (keep only letters) and split by whitespace
// We use a Regex to replace non-letters with spaces, then split by whitespace.
string cleanedContent = Regex.Replace(lowerContent, @"[^a-z\s]", " ");

// Split by whitespace, ignoring empty entries
string[] words = cleanedContent.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
var frequencyMap = new Dictionary<string, int>();

foreach (string word in words)
{
    if (word.Length == 0) continue;

    if (frequencyMap.ContainsKey(word))
    {
        frequencyMap[word]++;
    }
    else
    {
        frequencyMap[word] = 1;
    }
}

// Sort: by count descending, then by word alphabetically ascending
var sortedWords = frequencyMap
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

// Output results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}