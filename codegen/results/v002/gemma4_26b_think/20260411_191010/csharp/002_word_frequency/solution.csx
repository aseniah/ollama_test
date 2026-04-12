using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Define the path to the input file
string filePath = "input/text.txt";

// Check if the file exists before attempting to read
if (File.Exists(filePath))
{
    // Read the entire content of the file
    string content = File.ReadAllText(filePath);

    // Split the content into tokens using whitespace as a delimiter
    string[] tokens = content.Split(new char[] { ' ', '\r', '\n', '\t' }, StringSplitOptions.RemoveEmptyEntries);

    // Dictionary to store the frequency of each word
    var counts = new Dictionary<string, int>();

    foreach (string token in tokens)
    {
        // Use StringBuilder to construct the cleaned word (letters only, lowercase)
        StringBuilder sb = new StringBuilder();
        foreach (char c in token)
        {
            // Only keep characters that are letters
            if (char.IsLetter(c))
            {
                sb.Append(char.ToLowerInvariant(c));
            }
        }

        string word = sb.ToString();

        // If the resulting word is not empty, update its frequency in the dictionary
        if (!string.IsNullOrEmpty(word))
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
    }

    // Sort the results:
    // 1. By count (Value) in descending order
    // 2. By word (Key) in ascending alphabetical order for ties
    var sortedResults = counts.OrderByDescending(kvp => kvp.Value)
                              .ThenBy(kvp => kvp.Key);

    // Output each word and its count in the format "word: count"
    foreach (var kvp in sortedResults)
    {
        Console.WriteLine($"{kvp.Key}: {kvp.Value}");
    }
}