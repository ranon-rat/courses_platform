import { getFormData, handleResponse, request } from "./tools.js";

document.getElementById("signInForm")!.addEventListener("submit", async function (e) {
    let data: { [key: string]: any } = {};
    e.preventDefault();

    getFormData(e.target as HTMLFormElement).forEach(
        (value, key) => data[key] = value
    );

    if (handleResponse(await request('sign-in', data)))
        window.location.reload();
});

document.getElementById("signUpForm")!.addEventListener("submit", async function (e) {
    let data: { [key: string]: any } = {};
    e.preventDefault();

    getFormData(e.target as HTMLFormElement).forEach(
        (value, key) => data[key] = value
    );

    if (handleResponse(await request('sign-up', data)))
        window.location.reload();
});
