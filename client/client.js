const http = require('axios')

const website = "http://localhost:8080/"


// /*
// Tests the /users/create POST METHOD
// Also Tests the /users/search/:userid GET METHOD
// */
// http.post(website + 'users/create' , {
//         email: "omeralimalik96@gmail.com",
//         name: "Im secure",
//         password: "3717733579",
//         username: "three",
//     }).then(response => {
//         console.log("successfully created account getting it now")
//         http.get(website + "users/search/two").then(resp => {
//             console.log("get request successful!")
//             console.log(resp.data);
//         }).catch(err => {
//             console.log(err);
//         })
//     }).catch(err => {
//         console.log(err)
//     })

// /**
//  * Tests the /users/authenticate method
//  */
// http.post(website + "users/authenticate" , {
//     userId : "two",
//     password: "3717733578",
//     hashed: true,
// }).then(resp => {
//     console.log("authentication succesful!")
//     console.log("Results of authentication + "  + resp.data);
// }).catch(error => {
//     if (error.response) {
//         console.log(error.response.data);
//         console.log(error.response.status);
//         console.log(error.response.headers);
//     }
// })

// /**
//  * Tests the /users/update/:userid method
//  */

// http.post(website + "users/update/two" , {
//     email: "My email is updated!"
// }).then(resp => {
//     console.log(resp.data);
// }).catch(error => {
//     if (error.response) {
//         console.log(error.response.data);
//         console.log(error.response.status);
//         console.log(error.response.headers);
//     }
// })

http.get("http://localhost:8080/posts")
    .then(resp => console.log(resp.data))
    .catch(error => {
        console.log("err occured in Posts request")
        if (error.response) {
            console.log("data : " + error.response.data);
            console.log("status : " + error.response.status);
            console.log("headers : " + error.response.headers);
        } else {
            console.log(error)
        }
    })

console.log("Started Program!")