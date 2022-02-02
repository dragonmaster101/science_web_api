const http = require('axios')

const website = "http://localhost:8080/"



http.post(website + 'users/create' , {
        email: "omeralimalik96@gmail.com",
        name: "Im insecure",
        password: "Thisisnothashed",
        username: "one",
    }).then(response => {
        console.log("successfully created account getting it now")
        http.get(website + "users/search/one").then(resp => {
            console.log("get request successful!")
            console.log(resp.data);
        }).catch(err => {
            console.log(err);
        })
    }).catch(err => {
        console.log(err)
    })

console.log("Something is happening")