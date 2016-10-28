# Warnsdorff

Implements Warsdorff's Hueristic to find a [knight's tour](https://en.wikipedia.org/wiki/Knight%27s_tour) in an arbitrary chess board.

# Run
`go run main.go`

# Arguments
  ```
  -size=<size of board>
  -startX=<starting x position on the board>
  -startY=<starting y position on the board>
  ```
  
# Example
```go
>>> go run main.go -startX=1 -startY=1 -size=10
Found an open path

  23  62  51   2  25  74  57   4  27  30
  50   1  24  73  60   3  26  29  58   5
  63  22  61  52  75  66  59  56  31  28
  46  49  76  65  72  53 100  67   6  55
  21  64  47  92  97  78  71  54  99  32
  48  45  40  77  70  93  98  79  68   7
  41  20  89  96  91  82  69  94  33  80
  44  17  42  39  88  95  86  81   8  11
  19  38  15  90  83  36  13  10  85  34
  16  43  18  37  14  87  84  35  12   9
```
