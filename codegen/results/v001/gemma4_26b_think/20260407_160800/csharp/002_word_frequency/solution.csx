using System;
using System.IO;
using System.Linq;
using System.Text;

// Define the path to the input file
string filePath = "input/text.txt";

if (File.Exists(filePath))
{
    // Read the entire content of the file
    string text = File.ReadAllText(filePath);

    // Prepare a buffer to hold the processed characters
    char[] buffer = new char[text.Length];

    // Iterate through the text to convert to lowercase and strip anything that isn't a letter
    // Replacing non-letters with a space ensures that words separated by punctuation 
    // (e.g., "word1,word2") are correctly split into two distinct words.
    for (int i = 0; i < text.Length; i++)
    {
        if (char.IsLetter(text[i]))
        {
            buffer[i] = char.ToLowerInvariant(text[i]);
        }
        else
        {
            buffer[i] = ' ';
        }
    }

    // Create a new string from the processed characters
    string cleanedText = new string(buffer);

    // Split the string into words using whitespace as the delimiter
    // StringSplitOptions.RemoveEmptyEntries handles multiple spaces/newlines
    string[] words = cleanedText.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

    // Group the words, count their frequencies, and apply the sorting rules:
    // 1. Primary Sort: Count descending
    // 2. Secondary Sort: Alphabetical ascending
    var wordFrequencies = words
        .GroupBy(w => w)
        .Select(group => new 
        { 
            Word = group.Key, 
            Count = group.Count() 
        })
        .OrderByDescending(x => x.Count)
        .ThenBy(x => x.Word);

    // Output each result in the specified format
    foreach (var entry in wordFrequencies)
    {
        Console.WriteLine($"{entry.Word}: {entry.Count}");
    }
}