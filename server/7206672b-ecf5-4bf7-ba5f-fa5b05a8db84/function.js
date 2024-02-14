  function checkCeltAssociation() {

        var celt1 = process.env.MY_VARIABLE_INPUT_1; // Retrieve the value of CELT1 from environment variables

        var celt2 = process.env.MY_VARIABLE_OUTPUT; // Retrieve the value of CELT2 from environment variables

        console.log("celt1",celt1)

        console.log("celt2",celt2)



        if (celt1 === celt2) {

            process.exit(0)

        } else {

            process.exit(1)

        }

    }

    

    console.log(checkCeltAssociation());