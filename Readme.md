# Eurovent data to csv

A simple tool that queries data from Eurovent and uses its service to output a CSV.

## Installation

- (Beginners option): Download the latest binary for your operating system from
the Releases section
- (Slightly more advanced): Clone this repository to run it with golang instead

## Usage (with the binaries)

To search for `AC` (aircondition units) of type `AC1/A/S/R` (Air to air, split, reversible < 12KW)
for the brand of `Carrier` and create a CSV file named `data.csv` use:

    eurovent.exe -program "AC" -type "AC1/A/S/R" -brands "CARRIER" > data.csv

If you want to search in multiple brands, you can do so by using comma separated values such as:

    eurovent.exe -program "AC" -type "AC1/A/S/R" -brands "CARRIER,LG,MIDEA" > data.csv

Make sure you **do not have space between each Brand**. This will not work: `CARRIER, LG` (note the extra space).

The parameter values come directly from the [Eurovent Advanced Search webpage](https://www.eurovent-certification.com/en/advancedsearch/result?program=AC). As a result, the Brand names, Type, etc must match exactly
what is shown in that webpage.
