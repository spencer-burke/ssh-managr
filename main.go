package main

import (
"os"
"flag"
"fmt"
"term"
)

func main() {
//currently nothing yet
    text_ptr := flag.String("text", "", "Text to parse.")
    metric_ptr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
    unique_ptr := flag.Bool("unique", false, "Measure unique values of a metric.")
    flag.Parse()

    fmt.Printf("text_ptr: %s, metric_ptr: %s, unique_ptr: %t\n", *text_ptr, *metric_ptr, *unique_ptr)
}

