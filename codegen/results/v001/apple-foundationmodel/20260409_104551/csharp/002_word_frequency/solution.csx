using System;
using System.IO;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

// Function to read a file and count word frequencies
static void CountWordFrequencies()
{
    // Check if input file exists
    if (!File.Exists("input/text.txt"))
    {
        Console.WriteLine("Error: The file 'input/text.txt' does not exist.");
        return;
    }

    string filePath = "input/text.txt";
    using (StreamReader reader = new StreamReader(filePath))
    {
        string line;
        int totalWords = 0;
        int wordCounts = new int[256]; // Array to store frequency counts for each ASCII character (0-255)

        // Read lines from the file and count word frequencies
        while ((line = reader.ReadLine()) != null)
        {
            // Convert line to lowercase
            string lowerLine = line.ToLower();

            // Process each character in the line
            for (int i = 0; i < lowerLine.Length; i++)
            {
                char ch = lowerLine[i];
                // Check if character is a letter (ASCII 97-122 or 65-90)
                if (char.IsLetter(ch))
                {
                    wordCounts[ch - 'a']++;
                    wordCounts[ch - 'A']++;
                }
            }
            totalWords++;
        }

        // Convert wordCounts array to a frequency dictionary
        var wordFrequency = new Dictionary<char, int>();
        for (int i = 0; i < 256; i++)
        {
            if (wordCounts[i] > 0)
            {
                char wordChar = (char)(i + 'a');
                wordFrequency[wordChar] = wordCounts[i];
            }
        }

        // Sort words by frequency descending and alphabetically ascending
        var sortedWords = wordFrequency.OrderByDescending(kv => kv.Value).ThenBy(kv => kv.Key);

        // Output results
        foreach (var kvp in sortedWords)
        {
            Console.WriteLine($"{kvp.Key}: {kvp.Value}");
        }
    }
}

// Call the function to execute
CountWordFrequencies();