using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text;

// Define the path to the input file
string filePath = "input/text.txt";

// Check if the file exists before attempting to read
if (File.Exists(filePath))
{
    // Read all content from the file
    string text = File.ReadAllText(filePath);

    // Split the text into tokens using whitespace as the primary delimiter
    string[] tokens = text.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

    // Use a dictionary to store the frequency of each cleaned word
    var counts = new Dictionary<string, int>();

    foreach (string token in tokens)
    {
        // For each token, "strip all punctuation" by keeping only letters
        // and "convert all words to lowercase"
        StringBuilder sb = new StringBuilder();
        foreach (char c in token)
        {
            if (char.IsLetter(c))
            {
                sb.Append(char.ToLowerInvariant(c));
            }
        }

        string cleanedWord = sb.ToString();

        // If the token contained at least one letter, update the frequency map
        if (!string.IsNullOrEmpty(cleanedWord))
        {
            if (counts.ContainsKey(cleanedWord))
            {
                counts[cleanedWord]++;
            }
            else
            {
                counts[cleanedWord] = 1;
            }
        }
    }

    // Perform sorting:
    // 1. Sort by frequency (count) in descending order
    // 2. For ties, sort by the word itself in ascending alphabetical order
    var sortedResults = counts
        .OrderByDescending(kvp => kvp.Value)
        .ThenBy(kvp => kvp.Key);

    // Output the result in the format "word: count"
    foreach (var entry in sortedResults)
    {
        Console.WriteLine($"{entry.Key}: {entry.Value}");
    }
}