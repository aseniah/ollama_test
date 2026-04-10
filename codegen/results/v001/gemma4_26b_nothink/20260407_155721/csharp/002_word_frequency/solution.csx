using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;
using System.Text.RegularExpressions;

// Define the path to the input file
string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    // If the file doesn't exist, we exit silently as per requirements 
    // (no output allowed except the specific format).
    return;
}

// Read the entire content of the file
string content = File.ReadAllText(filePath);

// Use a dictionary to store word frequencies
var wordCounts = new Dictionary<string, int>();

// Regex to match only letters. 
// We split by anything that isn't a letter.
// This effectively strips punctuation and numbers.
string[] words = Regex.Split(content.ToLower(), @"[^a-z]+");

foreach (string word in words)
{
    // Regex.Split can result in empty strings if there are multiple non-alpha chars in a row
    if (!string.IsNullOrWhiteSpace(word))
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
}

// Sort logic:
// 1. Sort by Count Descending
// 2. Sort by Word Ascending (Alphabetical)
var sortedWords = wordCounts
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

// Output the results
foreach (var entry in sortedWords)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}