// Read and parse the CSV file
let csvLines = File.ReadAllText("input/data.csv") |> System.IO.File.ReadAllLines
let header = csvLines.[0] |> string.Split[,,,] |> System.Collections.Generic.List<string>.[^] // Skip header

// Parse data rows into objects
let dataRows = csvLines.[1..] |> Array.toList |> System.Collections.Generic.List<string>.[^] |> Array.ofSeq |> Array.map (fun rowStr ->
  let fields = rowStr |> string.Split(',') |> System.Collections.Generic.List<string>.[^]
  {
    Name = fields.[0] |> string
    Age = int.Parse fields.[1]
    Email = fields.[2] |> string
    Score = float.Parse fields.[3]
  }
) |> System.Collections.Generic.List<object>.[^] |> System.Collections.Generic.List<object>.ToList

// Convert list of objects to JSON array
let jsonOutput = System.Text.Json.JsonSerializer.Serialize(dataRows, System.Text.Json.JsonSerializerOptions { WriteIndented = false })

// Output JSON array to stdout
printfn "%s" jsonOutput