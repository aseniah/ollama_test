using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text;

// Define the path to the input file as specified
string filePath = "input/text.txt";

// Ensure the file exists before attempting to read it
if (File.Exists(filePath))
{
    // Read the entire content of the file
    string text = File.ReadAllText(filePath);

    // Create a StringBuilder to process the text
    // We iterate through every character to implement the rule:
    // "Strip all punctuation (keep only letters)"
    // We also convert characters to lowercase during this step.
    StringBuilder sb = new StringBuilder();
    foreach (char c in text)
    {
        if (char.IsLetter(c))
        {
            // If it's a letter, append its lowercase version
            sb.Append(char.ToLowerInvariant(c));
        }
        else
        {
            // If it's punctuation or whitespace, replace it with a space
            // This ensures that words separated by punctuation (like "times,it") 
            // are correctly split into separate words.
            sb.Append(' ');
        }
    }

    // Split the processed string into an array of words, 
    // removing any empty strings caused by multiple consecutive spaces.
    string[] words = sb.ToString().Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

    // Use a dictionary to track the frequency of each unique word
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

    // Apply the sorting rules:
    // 1. Primary sort: Count descending (OrderByDescending)
    // 2. Secondary sort: Alphabetical ascending (ThenBy)
    var sortedResults = counts.OrderByDescending(kvp => kvp.Value)
                              .ThenBy(kvp => kvp.Key);

    // Print the results to stdout in the format "word: count"
    foreach (var entry in sortedResults)
    {
        Console.WriteLine($"{entry.Key}: {entry.Value}");
    }
}