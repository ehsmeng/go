package load_csv

// (c) Marcus Engene myfirstname@mylastname.se BSD license. Free to use in commercial projects

// convenience package to load CSV into float64 array of arrays

import (
    "os"
    "io"
    "encoding/csv"
    "strconv"
)

// nbr_of_rows, nbr_of_cols
func Count_csv_lines_and_fields(filename string) (error, int, int) {
    nbr_of_rows := 0
    nbr_of_cols := -1

    csvfile, err := os.Open(filename)

    if err != nil {
        return err, -1, -1
    }

    defer csvfile.Close()

    reader := csv.NewReader(csvfile)
    for {
        record, err := reader.Read()
        // end-of-file is fitted into err
        if err == io.EOF {
            break
        } else if err != nil {
            return err, -1, -1
        }

        if 0 == len(record) {
            continue
        }

        nbr_of_rows += 1

        if -1 == nbr_of_cols {
            nbr_of_cols = len(record)
        }

        // column mismatch is part of csv read thing
    }

    return nil, nbr_of_rows, nbr_of_cols
}


func Readcsv_float64(filename string) (error, int, int, [][]float64) {
    err, nbr_of_rows, nbr_of_cols := Count_csv_lines_and_fields(filename)
    if nil != err {
        return err, -1, -1, nil
    }

    csv_array := make([][]float64, nbr_of_rows)

    csvfile, err := os.Open(filename)

    if err != nil {
        return err, -1, -1, nil
    }

    defer csvfile.Close()

    reader := csv.NewReader(csvfile)
    atrow := 0
    for {
        record, err := reader.Read()
        // end-of-file is fitted into err
        if err == io.EOF {
            break
        } else if err != nil {
            return err, -1, -1, nil
        }

        if 0 == len(record) {
            continue
        }

        row := make([]float64, nbr_of_cols)
        for i:=0; i<nbr_of_cols; i++ {
            v, err := strconv.ParseFloat(record[i], 64)
            if nil != err {
                return nil, -1, -1, nil
            }
            row[i] = v
        }

        csv_array[atrow] = row
        atrow += 1
    }

    return nil, nbr_of_rows, nbr_of_cols, csv_array
}
