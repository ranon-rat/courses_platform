import { getFormData, handleResponse, request } from "./tools.js";

document.getElementById("new-post")!.addEventListener("submit", async function (e) {
    let data: { [key: string]: any } = {};
    e.preventDefault();

    getFormData(e.target as HTMLFormElement).forEach(
        (value, key) => data[key] = value
    );

    data["content"] = await data["content"].text()

    if (handleResponse(await request('new-post', data)))
        window.location.reload();
});
