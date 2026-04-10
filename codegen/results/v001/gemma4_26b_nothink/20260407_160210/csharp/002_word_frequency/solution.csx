using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using System.Collections.Generic;

// Define the file path
string filePath = "input/text.txt";

try
{
    // Check if file exists to avoid runtime crash
    if (!File.Exists(filePath))
    {
        return;
    }

    // Read all text from the file
    string text = File.ReadAllText(filePath);

    // Use Regex to find all words (sequences of letters)
    // This naturally handles lowercase conversion and stripping punctuation
    var matches = Regex.Matches(text.ToLowerInvariant(), @"[a-z]+");

    // Group matches by value and count them
    var wordCounts = matches
        .Cast<Match>()
        .Select(m => m.Value)
        .GroupBy(word => word)
        .Select(group => new 
        { 
            Word = group.Key, 
            Count = group.Count() 
        })
        // Sort by Count descending, then by Word ascending
        .OrderByDescending(x => x.Count)
        .ThenBy(x => x.Word)
        .ToList();

    // Output the results in the required format
    foreach (var item in wordCounts)
    {
        Console.WriteLine($"{item.Word}: {item.Count}");
    }
}
catch (Exception)
{
    // Silent exit as per "Do not output anything else" rule if errors occur
}