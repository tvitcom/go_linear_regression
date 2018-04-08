package main

import (
        "encoding/csv"
        "fmt"
        "log"
        "os"
        "strconv"
//        "github.com/gonum/plot"
//        "github.com/gonum/plot/plotter"
//        "github.com/gonum/plot/vg"
        chart "gonum.org/v1/plot"
        chartplotter "gonum.org/v1/plot/plotter"
        vg "gonum.org/v1/plot/vg"
//        "image/color"
)

func main() {
    //we open csv file from the disk
    f, err := os.Open("./data/kc_house_data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    //we will create a new csv reader specifying the number of collumns 
    //it has
    salesData := csv.NewReader(f)
    salesData.FieldsPerRecord = 21

    //we read all the records:
    records, err := salesData.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    header := records[0]
    //by slicing the records we skip the header
    records=records[1:]

    //we iterate over all records
    //and keep track of all the gathered values
    //for each column
    columnsValues := map[int]chartplotter.Values{}
    for i, record := range records {
        //we want one histogram per column,
        //so we will iterate over all the columns we have
        //and gather the date or each in a separate value set
        //in columnValues
        //we are skipping the ID  column and the Date
        //so we start on  index 2.
        for c := 2; c < salesData.FieldsPerRecord; c++ {
            if _, found := columnsValues[c]; !found {
               columnsValues[c] =  make(chartplotter.Values, len(records))
            }
            //we parse each close value and add it to our set
            floatVal, err := strconv.ParseFloat(record[c],64)
            if err != nil {
                log.Fatal(err)
            }
            columnsValues[c][i] = floatVal
        }
    }

    //once we have all the data, we draw each graph
    for c, values := range columnsValues {
        //create new plot
        p, err := chart.New()
        if err != nil {
            log.Fatal(err)
        }
        p.Title.Text = fmt.Sprintf("Histogram of %s", header[c])

        //create a new normalize histogram
        //and add it to the plot
        h, err := chartplotter.NewHist(values, 16)
        if err != nil {
            log.Fatal(err)
        }
        h.Normalize(1)
        p.Add(h)

        //save the plot to a png file
        if err := p.Save(
            10*vg.Centimeter,
            10*vg.Centimeter,
            fmt.Sprintf("./graphs/%s_hist.png", header[c]),
        ); err != nil {
            log.Fatal(err)
        }
    }
}
