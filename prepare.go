package main

import (
    "encoding/csv"
    "log"
    "math/rand"
    "os"
)

func main() {
    //we nedd to open csv file:
    f, err := os.Open("./data/kc_house_data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    //we create a new csv render specifying the number of columns it has
    salesData := csv.NewReader(f)
    salesData.FieldsPerRecord = 21
    
    //we read all the record
    records, err := salesData.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    //save the header
    header := records[0]

    // we have to shuffle the dataset befor esplitting 
    // to avoid having ordered data
    //if the data is ordered, th data in the train set and the one in 
    //the test set, can have different behavior
    shuffled := make([][]string, len(records)-1)
    perm := rand.Perm(len(records)-1)
    for i, v := range perm {
        shuffled[v] = records[i+1]
    }

    //split the training set 
    trainingIdx := (len(shuffled)) * 4 / 5
    trainingSet := shuffled[1:trainingIdx+1]
    
    //split the testing set 
    testingSet := shuffled[trainingIdx+1:]

    // we write the splitted set in the separated files
    sets := map[string][][]string{
        "./data/training.csv": trainingSet,
        "./data/testing.csv": testingSet,
    }

    for fn, dataset := range sets {
        f, err :=os.Create(fn)
        if err != nil {
            log.Fatal(err)
        }

        defer f.Close();

        out := csv.NewWriter(f)
        if err := out.Write(header); err != nil {
            log.Fatal(err)
        }

        if err := out.WriteAll(dataset); err != nil {
            log.Fatal(err)
        }
        out.Flush()
    }
}

