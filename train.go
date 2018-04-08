package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "log"
    "strconv"

    "github.com/sajari/regression"
)

func main() {
    //we open csv file
    f, err := os.Open("./data/training.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    //we need create a new  csv reader specifying
    //the number of columns it has
    salesData := csv.NewReader(f)
    salesData.FieldsPerRecord = 21

    //we read all the records
    records, err := salesData.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    //in this case we are going to try and model our house price (y)
    //by the egrade feature
    var r regression.Regression
    r.SetObserved("Price")
    r.SetVar(0,"Grade")

    //loop of records in the CSV, adding the training data to regression value.
    for i, record := range records {

        //skip the header
        if i==0 {
            continue
        }

        //parce the house price, "y".
        price, err := strconv.ParseFloat(records[i][2], 64)
        if err != nil {
            log.Fatal(err)
        }

        //Parce the grade value
        grade, err := strconv.ParseFloat(record[11], 64)
        if err != nil {
            log.Fatal(err)
        }

        //add this points to the regression value.
        r.Train(regression.DataPoint(price, []float64{grade}))
    }

    //Train /fit the regression model.
    r.Run()
    //Output the trained model parameters.
    fmt.Printf("\nRegression formula:\n%v\n\n",r.Formula)
}
