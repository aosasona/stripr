/* require the http module
create a server instance
*/
const http = require("http");
const server = http.createServer()

// respond to all requests
server.on("request", (req, res) => {
    res.end("Hello World!")
})

/* start the server on port 3000 */
server.listen(3000, () => {
    console.log("Server is listening on port 3000"); // show this message in the terminal once the server starts
})