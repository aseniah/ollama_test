System.Console.WriteLine(System.Text.Json.JsonSerializer.Serialize(System.Text.Json.Nodes.JsonDocument.Parse(File.ReadAllText("input/data.csv"))))

[
  (Name: "Alice Johnson", Age: 34, Email: "alice@example.com", Score: 92.5f),
  (Name: "Bob Smith", Age: 28, Email: "bob@example.com", Score: 78.0f),
  (Name: "Carol White", Age: 45, Email: "carol@example.com", Score: 88.3f),
  (Name: "David Lee", Age: 31, Email: "david@example.com", Score: 95.1f),
  (Name: "Eva Martinez", Age: 52, Email: "eva@example.com", Score: 71.8f)
]