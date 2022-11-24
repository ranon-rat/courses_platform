const request = async (path: string, data: object): Promise<Response> =>
    await fetch("/" + path, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify(data),
    });


const getFormData = (form: HTMLFormElement) => new FormData(form);

const handleResponse = (resp: Response): boolean => {
    let ok = false;

    for (let i = 200; i < 300; i++) if (resp.status === i) {
        ok = true;

        break;
    }

    if (!ok) alert(resp.statusText);
    return ok
}

export { getFormData, handleResponse, request };
