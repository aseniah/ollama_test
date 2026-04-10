using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

// Define the path to the input file
string filePath = "input/text.txt";

// Check if the file exists before attempting to read
if (File.Exists(filePath))
{
    // Read the entire content of the file
    string content = File.ReadAllText(filePath);

    // Split the content into segments using whitespace as the delimiter
    string[] segments = content.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

    // Dictionary to store the word frequencies
    var wordCounts = new Dictionary<string, int>();

    foreach (string segment in segments)
    {
        // Filter the segment to keep only letters and convert to lowercase
        string cleanedWord = new string(segment.Where(c => char.IsLetter(c)).ToArray()).ToLower();

        // Only process if the word is not empty after stripping punctuation/numbers
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

    // Sort by frequency descending, then by word alphabetically ascending
    var sortedResults = wordCounts
        .OrderByDescending(kvp => kvp.Value)
        .ThenBy(kvp => kvp.Key);

    // Output the results in the required format
    foreach (var entry in sortedResults)
    {
        Console.WriteLine($"{entry.Key}: {entry.Value}");
    }
}