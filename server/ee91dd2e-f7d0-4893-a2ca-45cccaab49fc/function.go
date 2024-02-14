package main



import (

    "os"

)



func checkCeltAssociation() {

    celt1 := os.Getenv("MY_VARIABLE_INPUT_1") // קבל את הערך של CELT1 מהמשתנים הסביבתיים

    celt2 := os.Getenv("MY_VARIABLE_OUTPUT") // קבל את הערך של CELT2 מהמשתנים הסביבתיים



    if celt1 == celt2 {

         os.Exit(0) // מקושר

    } else {

         os.Exit(1) // לא מקושר

    }

}



func main() {

    checkCeltAssociation()

}

