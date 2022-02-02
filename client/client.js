const http = require('axios')

const website = "http://localhost:8080/"



http.post(website + 'users/create' , {
        email: "omeralimalik96@gmail.com",
        name: "Im secure",
        password: "3717733579",
        username: "three",
    }).then(response => {
        console.log("successfully created account getting it now")
        http.get(website + "users/search/two").then(resp => {
            console.log("get request successful!")
            console.log(resp.data);
        }).catch(err => {
            console.log(err);
        })
    }).catch(err => {
        console.log(err)
    })

http.post(website + "users/authenticate" , {
    userId : "two",
    password: "3717733578",
    hashed: true,
}).then(resp => {
    console.log("authentication succesful!")
    console.log("Results of authentication + "  + resp.data);
}).catch(error => {
    if (error.response) {
        console.log(error.response.data);
        console.log(error.response.status);
        console.log(error.response.headers);
    }
})

console.log("Started Program!")