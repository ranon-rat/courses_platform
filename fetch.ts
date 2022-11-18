await fetch("http://localhost:8080/sign-up", {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
        email: "test@gmail.com",
        password: "1234",
        username: "test",
        privileges: 1,
    }),
})