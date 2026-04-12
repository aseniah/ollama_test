using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;

// Define the path to the input file
string filePath = "input/text.txt";

// Check if the file exists before attempting to read it
if (File.Exists(filePath))
{
    // Read the entire content of the file
    string content = File.ReadAllText(filePath);

    // Split the content into raw tokens using whitespace (spaces, tabs, newlines) as delimiters
    // This handles the initial separation of words
    string[] rawTokens = content.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

    // Dictionary to store the frequency of each processed word
    var wordFrequencies = new Dictionary<string, int>();

    foreach (string token in rawTokens)
    {
        // Process each token to "Strip all punctuation (keep only letters)"
        // We use a StringBuilder for efficient string construction
        var sb = new StringBuilder();
        foreach (char c in token)
        {
            // Check if the character is a letter (as per the rule "keep only letters")
            if (char.IsLetter(c))
            {
                // Convert character to lowercase as per the rule "Convert all words to lowercase"
                sb.Append(char.ToLowerInvariant(c));
            }
        }

        string cleanedWord = sb.ToString();

        // Only add to the dictionary if the word is not empty after stripping punctuation
        if (!string.IsNullOrEmpty(cleanedWord))
        {
            if (wordFrequencies.ContainsKey(cleanedWord))
            {
                wordFrequencies[cleanedWord]++;
            }
            else
            {
                wordFrequencies[cleanedWord] = 1;
            }
        }
    }

    // Sort the results: 
    // 1. By count descending (OrderByDescending)
    // 2. For ties, sort by word alphabetically ascending (ThenBy)
    var sortedResults = wordFrequencies
        .OrderByDescending(kvp => kvp.Value)
        .ThenBy(kvp => kvp.Key);

    // Output the final results in the format: word: count
    foreach (var entry in sortedResults)
    {
        Console.WriteLine($"{entry.Key}: {entry.Value}");
    }
}